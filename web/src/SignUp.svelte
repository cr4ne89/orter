<script>

  let username = '';
  let password = '';

  const signUp = async () => {
    try {
      const response = await fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({ username, password }),
      });

      if (response.ok) {
        alert('User created successfully');
        // Redirect to the main page after sign-up
        window.location.href = '/'
      } else {
        alert('Error creating user, please try again.');
      }
    } catch (error) {
      console.error('SignUp error:', error);
      alert('An error occurred while signing up. Please try again later.');
    }
  };
</script>

<div class="signup-container">
  <h2>Sign Up</h2>
  <input type="text" bind:value={username} placeholder="Username" />
  <input type="password" bind:value={password} placeholder="Password" />
  <button on:click={signUp}>Sign Up</button>
</div>

<style>
  .signup-container {
    max-width: 400px;
    margin: auto;
    text-align: center;
  }
  input {
    display: block;
    margin: 10px auto;
    padding: 8px;
    width: 80%;
  }
  button {
    padding: 10px 20px;
  }
</style>
