<script>
  import { userSession } from './stores';

  const logout = () => {
    userSession.set(null);
    localStorage.removeItem('username');
    alert('Logout successful');
    window.location.href = '/';
  };
</script>

<div class="logout-container">
  <h2>Logout</h2>
  <p>Are you sure you want to logout?</p>
  <button on:click={logout}>Yes, Logout</button>
</div>

<style>
  .logout-container {
    text-align: center;
  }
  button {
    padding: 10px 20px;
    margin: 10px;
  }
</style>