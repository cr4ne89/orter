package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

var (
	db    *sql.DB
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
)

func main() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to API")
	}).Methods("GET")
	r.HandleFunc("/signup", handleSignUp).Methods("POST")
	r.HandleFunc("/login", handleLogin).Methods("POST")
	r.HandleFunc("/logout", handleLogout).Methods("POST")
	r.HandleFunc("/daily-log", handleDailyLog).Methods("POST", "GET")
	r.HandleFunc("/items", handleItems).Methods("POST")

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
	}).Handler(r)

	port := "8080"
	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User created successfully")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var hashedPassword string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username=$1", username).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["username"] = username
	session.Save(r, w)

	fmt.Fprintln(w, "Login successful")
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// セッションを無効化
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)

	fmt.Fprintln(w, "Logout successful")
}

func handleDailyLog(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	username := session.Values["username"].(string)
	var userID string
	err := db.QueryRow("SELECT user_id FROM users WHERE username=$1", username).Scan(&userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodPost:
		date := r.FormValue("date")
		items := r.FormValue("items")

		_, err := db.Exec("INSERT INTO daily_logs (user_id, date) VALUES ($1, $2)", userID, date)
		if err != nil {
			http.Error(w, "Could not create daily log", http.StatusInternalServerError)
			return
		}

		var itemList []map[string]interface{}
		if err := json.Unmarshal([]byte(items), &itemList); err != nil {
			http.Error(w, "Invalid items format", http.StatusBadRequest)
			return
		}

		for _, item := range itemList {
			itemID := item["item_id"].(string)
			data := item["data"].(map[string]interface{})
			_, err := db.Exec("INSERT INTO daily_log_items (daily_log_id, item_id, data) VALUES ($1, $2, $3)", userID, itemID, data)
			if err != nil {
				http.Error(w, "Could not add item to daily log", http.StatusInternalServerError)
				return
			}
		}

		fmt.Fprintln(w, "Daily log created successfully")

	case http.MethodGet:
		date := r.URL.Query().Get("date")
		rows, err := db.Query("SELECT day_id, date FROM daily_logs WHERE user_id=$1 AND date=$2", userID, date)
		if err != nil {
			http.Error(w, "Could not retrieve daily logs", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var logs []map[string]interface{}
		for rows.Next() {
			var dayID string
			var logDate time.Time
			if err := rows.Scan(&dayID, &logDate); err != nil {
				http.Error(w, "Error reading daily logs", http.StatusInternalServerError)
				return
			}

			logs = append(logs, map[string]interface{}{
				"day_id": dayID,
				"date":   logDate,
			})
		}

		jsonResponse, err := json.Marshal(logs)
		if err != nil {
			http.Error(w, "Error encoding daily logs to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleItems(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	username := session.Values["username"].(string)
	var userID string
	err := db.QueryRow("SELECT user_id FROM users WHERE username=$1", username).Scan(&userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		itemType := r.FormValue("item_type")

		_, err := db.Exec("INSERT INTO items (user_id, name, description, item_type) VALUES ($1, $2, $3, $4)", userID, name, description, itemType)
		if err != nil {
			http.Error(w, "Could not create item", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Item created successfully")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
