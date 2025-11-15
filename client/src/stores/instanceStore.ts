import { writable } from 'svelte/store';
import type { Writable } from 'svelte/store';
import * as api from '../lib/api';
import type { Instance, NewInstanceRequest } from '../lib/types';

interface InstanceStore extends Writable<Instance[]> {
  load: (user: string) => Promise<void>;
  create: (user: string, data: NewInstanceRequest) => Promise<Instance>;
  save: (user: string, instance: Instance) => Promise<void>;
}

function createInstanceStore(): InstanceStore {
  const { subscribe, set, update } = writable<Instance[]>([]);

  return {
    subscribe,
    set,
    update,
    
    load: async (user: string) => {
      try {
        const instances = await api.getInstances(user);
        set(instances);
      } catch (error) {
        console.error('Failed to load instances:', error);
        throw error;
      }
    },
    
    create: async (user: string, data: NewInstanceRequest) => {
      try {
        const newInstance = await api.createInstance(user, data);
        update(instances => [...instances, newInstance]);
        return newInstance;
      } catch (error) {
        console.error('Failed to create instance:', error);
        throw error;
      }
    },
    
    save: async (user: string, instance: Instance) => {
      try {
        await api.saveInstance(user, instance);
        update(instances => 
          instances.map(i => i._id === instance._id ? instance : i)
        );
      } catch (error) {
        console.error('Failed to save instance:', error);
        throw error;
      }
    }
  };
}

export const instances = createInstanceStore();
