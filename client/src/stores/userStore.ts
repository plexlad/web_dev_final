import { writable } from "svelte/store";
import type { Writable } from "svelte/store";

function getStoredUser(): string | null {
  if (typeof window === "undefined") return null;
  const stored = localStorage.getItem('username');
  return stored || null;
}

interface UserStore extends Writable<string | null> {
  setUser: (username: string) => void
  clearUser: () => void;
}

function createUserStore(): UserStore {
  const { subscribe, set, update } = writable<string | null>(getStoredUser());

  return {
    subscribe,
    set,
    update,
    setUser: (username: string) => {
      localStorage.setItem('username', username);
      set(username);
    },
    clearUser: () => {
      localStorage.removeItem('username');
      set(null);
    }
  }
}

export const currentUser = createUserStore();
