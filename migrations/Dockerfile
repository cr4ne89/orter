FROM golang:1.23

WORKDIR /app

RUN go install github.com/rubenv/sql-migrate/...@latest

COPY ./migrations /migrations

ENTRYPOINT ["sql-migrate", "up", "-config=/app/dbconfig.yml"]