<script>
  import { writable } from 'svelte/store';
  import Daily from './Daily.svelte';
  import Setting from './Setting.svelte';
  import Logout from './Logout.svelte';

  let showDaily = false;
  let showSetting = false;
  let showLogout = false;
  let selectedDate = '';
</script>

<div>
  <h2>Dashboard</h2>
  <div class="calendar">
    {#each Array(30) as _, index}
      <button on:click={() => { showDaily = true; selectedDate = `2024-11-${index + 1}`; }}>
        2024-11-{index + 1}
      </button>
    {/each}
  </div>
  <button on:click={() => showSetting = true} class="setting-button">Settings</button>
  <button on:click={() => showLogout = true} class="logout-button">Logout</button>

  {#if showDaily}
    <div class="modal">
      <Daily date={selectedDate} />
      <button on:click={() => showDaily = false}>Close</button>
    </div>
  {/if}

  {#if showSetting}
    <div class="modal">
      <Setting />
      <button on:click={() => showSetting = false}>Close</button>
    </div>
  {/if}

  {#if showLogout}
    <div class="modal">
      <Logout />
      <button on:click={() => showLogout = false}>Cancel</button>
    </div>
  {/if}
</div>

<style>
  .calendar {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    gap: 10px;
  }
  .setting-button {
    position: absolute;
    top: 10px;
    right: 60px;
  }
  .logout-button {
    position: absolute;
    top: 10px;
    right: 10px;
  }
  .modal {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: white;
    padding: 20px;
    box-shadow: 0px 0px 10px rgba(0,0,0,0.5);
    width: 300px;
  }
  button {
    padding: 10px;
    margin: 5px;
  }
</style>