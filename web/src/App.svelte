<script>
  import { onMount } from 'svelte';
  import { userSession } from './stores';
  import Login from './Login.svelte';
  import Dashboard from './Dashboard.svelte';

  onMount(() => {
    const username = localStorage.getItem('username');
    if (username) {
      userSession.set(username);
    }
  });
</script>

{#if $userSession}
  <Dashboard />
{:else}
  <Login />
{/if}
