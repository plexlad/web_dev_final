<script lang="ts">
  import { currentUser } from "../stores/userStore";

  let username = "";

  function handleLogin() {
    if (username.trim()) {
      currentUser.setUser(username.trim());
    }
  }

  function handleLogout() {
    currentUser.clearUser();
    username = "";
  }
</script>

{#if $currentUser}
  <div class="user-info">
    <p>Logged in as: <strong>{$currentUser}</strong></p>
    <button on:click={handleLogout}>Logout</button>
  </div>
{:else}
  <div class="login-form">
    <h2>Enter Your Username</h2>
    <form on:submit|preventDefault={handleLogin}>
      <input
        type="text"
        bind:value={username}
        placeholder="Username"
        required
      />
      <button type="submit">Continue</button>
    </form>
  </div>
{/if}

<style>
  .login-form {
    max-width: 400px;
    margin: 2rem auto;
    padding: 2rem;
    border: 1px solid #ddd;
    border-radius: 8px;
    background: white;
  }
  
  h2 {
    margin-top: 0;
    color: #333;
  }
  
  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  
  input {
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 1rem;
  }
  
  button {
    padding: 0.75rem;
    background: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 1rem;
    cursor: pointer;
  }
  
  button:hover {
    background: #45a049;
  }
  
  .user-info {
    padding: 1rem;
    background: #f0f0f0;
    border-radius: 4px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .user-info button {
    background: #f44336;
    padding: 0.5rem 1rem;
  }
  
  .user-info button:hover {
    background: #da190b;
  }
</style>
