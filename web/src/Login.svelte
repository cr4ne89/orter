<script>
  import { userSession } from './stores';
  import SignUp from './SignUp.svelte';

  let username = '';
  let password = '';
  let showSignUp = false;

  const login = async () => {
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({ username, password }),
      });

      if (response.ok) {
        userSession.set(username);
        localStorage.setItem('username', username);
        alert('Login successful');
      } else {
        alert('Invalid login, please try again.');
      }
    } catch (error) {
      console.error('Login error:', error);
      alert('An error occurred while logging in. Please try again later.');
    }
  };
</script>

{#if showSignUp}
  <SignUp />
{:else}
  <div class="login-container">
    <h2>Login</h2>
    <input type="text" bind:value={username} placeholder="Username" />
    <input type="password" bind:value={password} placeholder="Password" />
    <button on:click={login}>Login</button>
    <button on:click={() => showSignUp = true}>Create an Account</button>
  </div>
{/if}

<style>
  .login-container {
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
