<!-- src/App.svelte -->
<script lang="ts">
  import { currentUser } from './stores/userStore';
  import UserLogin from './components/UserLogin.svelte';
  import SchemaList from './components/SchemaList.svelte';
  import InstanceList from './components/InstanceList.svelte';
</script>

<main>
  <header>
    <h1>gardi</h1>
    <UserLogin />
  </header>

  {#if $currentUser}
    <div class="content">
      <section>
        <h2>Schemas</h2>
        <SchemaList user={$currentUser} />
      </section>
      
      <section>
        <h2>Instances</h2>
        <InstanceList user={$currentUser} />
      </section>
    </div>
  {:else}
    <div class="welcome">
      <p>Please enter your username to continue</p>
    </div>
  {/if}
</main>

<style>
  main {
    max-width: 1200px;
    margin: 0 auto;
    padding: 1rem;
  }
  
  header {
    margin-bottom: 2rem;
  }
  
  h1 {
    text-align: center;
    color: #333;
  }
  
  .content {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 2rem;
  }
  
  section {
    padding: 1rem;
    border: 1px solid #ddd;
    border-radius: 8px;
  }
  
  .welcome {
    text-align: center;
    padding: 2rem;
    color: #666;
  }
</style>
