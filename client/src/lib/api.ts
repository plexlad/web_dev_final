import type { Schema, Instance, NewSchemaRequest, NewInstanceRequest, ApiError } from './types';

const API_BASE = 'http://localhost:5499';

async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const error: ApiError = await response.json();
    throw new Error(error.error || 'API request failed');
  }
  return response.json();
}

// Schema API
export async function getSchemas(user: string): Promise<Schema[]> {
  const response = await fetch(`${API_BASE}/${user}/schemas`);
  return handleResponse<Schema[]>(response);
}

export async function createSchema(user: string, data: NewSchemaRequest): Promise<Schema> {
  const response = await fetch(`${API_BASE}/${user}/schemas/new`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  });
  return handleResponse<Schema>(response);
}

export async function saveSchema(user: string, schema: Schema): Promise<string> {
  const response = await fetch(`${API_BASE}/${user}/schemas/save`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(schema)
  });
  return response.text();
}

// Instance API
export async function getInstances(user: string): Promise<Instance[]> {
  const response = await fetch(`${API_BASE}/${user}/instances`);
  return handleResponse<Instance[]>(response);
}

export async function createInstance(user: string, data: NewInstanceRequest): Promise<Instance> {
  const response = await fetch(`${API_BASE}/${user}/instances/new`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  });
  return handleResponse<Instance>(response);
}

export async function saveInstance(user: string, instance: Instance): Promise<string> {
  const response = await fetch(`${API_BASE}/${user}/instances/save`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(instance)
  });
  return response.text();
}
