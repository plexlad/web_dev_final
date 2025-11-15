import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';
import * as api from '../lib/api';
import type { Schema, NewSchemaRequest } from '../lib/types';

interface SchemaStore extends Writable<Schema[]> {
  load: (user: string) => Promise<void>;
  create: (user: string, data: NewSchemaRequest) => Promise<Schema>;
  save: (user: string, schema: Schema) => Promise<void>;
}

function createSchemaStore(): SchemaStore {
  const { subscribe, set, update } = writable<Schema[]>([]);

  return {
    subscribe,
    set,
    update,
    
    load: async (user: string) => {
      try {
        const schemas = await api.getSchemas(user);
        set(schemas);
      } catch (error) {
        console.error('Failed to load schemas:', error);
        throw error;
      }
    },
    
    create: async (user: string, data: NewSchemaRequest) => {
      try {
        const newSchema = await api.createSchema(user, data);
        update(schemas => [...schemas, newSchema]);
        return newSchema;
      } catch (error) {
        console.error('Failed to create schema:', error);
        throw error;
      }
    },
    
    save: async (user: string, schema: Schema) => {
      try {
        await api.saveSchema(user, schema);
        update(schemas => 
          schemas.map(s => s.id === schema.id ? schema : s)
        );
      } catch (error) {
        console.error('Failed to save schema:', error);
        throw error;
      }
    }
  };
}

export const schemas = createSchemaStore();
