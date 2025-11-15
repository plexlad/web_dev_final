<script lang="ts">
  import { onMount } from 'svelte';
  import { schemas } from '../stores/schemaStore';
  import type { Instance, NewInstanceRequest } from '../lib/types';
  import * as api from '../lib/api';
  
  export let user: string;
  
  let instances: Instance[] = [];
  let loading = true;
  let error = '';
  let showForm = false;
  let newInstance: NewInstanceRequest = {
    name: '',
    description: '',
    schema_id: ''
  };
  
  onMount(async () => {
    try {
      // Load schemas first (for the dropdown)
      await schemas.load(user);
      // Load instances
      instances = await api.getInstances(user);
      loading = false;
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to load instances';
      loading = false;
    }
  });
  
  async function handleCreate() {
    try {
      const created = await api.createInstance(user, newInstance);
      instances = [...instances, created];
      newInstance = { name: '', description: '', schema_id: '' };
      showForm = false;
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to create instance';
    }
  }
  
  function getSchemaName(schemaId: string): string {
    const schema = $schemas.find(s => s.id === schemaId);
    return schema ? schema.name : 'Unknown Schema';
  }
</script>

<div class="instance-list">
  <div class="header">
    <h3>Instances</h3>
    <button on:click={() => showForm = !showForm}>
      {showForm ? 'Cancel' : '+ New Instance'}
    </button>
  </div>
  
  {#if error}
    <div class="error">{error}</div>
  {/if}
  
  {#if showForm}
    <form on:submit|preventDefault={handleCreate} class="create-form">
      <input 
        type="text" 
        bind:value={newInstance.name}
        placeholder="Instance name"
        required
      />
      <textarea 
        bind:value={newInstance.description}
        placeholder="Description"
        rows="3"
      ></textarea>
      <select bind:value={newInstance.schema_id} required>
        <option value="">Select a schema...</option>
        {#each $schemas as schema}
          <option value={schema._id}>{schema.name}</option>
        {/each}
      </select>
      <button type="submit">Create</button>
    </form>
  {/if}
  
  {#if loading}
    <div class="loading">Loading instances...</div>
  {:else if instances.length === 0}
    <div class="empty">No instances yet. Create one to get started!</div>
  {:else}
    <div class="grid">
      {#each instances as instance}
        <div class="card">
          <h4>{instance.name}</h4>
          <p>{instance.description}</p>
          <div class="meta">
            <small>Schema: {getSchemaName(instance.schema_id)}</small>
            <small>Created: {new Date(instance.created_at).toLocaleDateString()}</small>
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>

<style>
  .instance-list {
    width: 100%;
  }
  
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
  }
  
  h3 {
    margin: 0;
  }
  
  button {
    padding: 0.5rem 1rem;
    background: #2196F3;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 0.9rem;
  }
  
  button:hover {
    background: #1976D2;
  }
  
  .create-form {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
    padding: 1rem;
    background: #f9f9f9;
    border-radius: 4px;
    margin-bottom: 1rem;
  }
  
  input, textarea, select {
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 0.9rem;
    font-family: inherit;
  }
  
  .error {
    padding: 0.75rem;
    background: #ffebee;
    color: #c62828;
    border-radius: 4px;
    margin-bottom: 1rem;
  }
  
  .loading, .empty {
    text-align: center;
    padding: 2rem;
    color: #666;
  }
  
  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 1rem;
  }
  
  .card {
    padding: 1rem;
    border: 1px solid #ddd;
    border-radius: 4px;
    background: white;
    transition: box-shadow 0.2s;
  }
  
  .card:hover {
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
  }
  
  .card h4 {
    margin: 0 0 0.5rem 0;
    color: #333;
  }
  
  .card p {
    margin: 0 0 1rem 0;
    color: #666;
    font-size: 0.9rem;
  }
  
  .meta {
    display: flex;
    justify-content: space-between;
    font-size: 0.8rem;
    color: #999;
  }
</style>
