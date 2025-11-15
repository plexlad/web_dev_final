export interface Schema {
  _id: string;
  version: number;
  user_version: number;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}

export interface Instance {
  _id: string,
  schema_id: string;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}

export interface NewSchemaRequest {
  name: string;
  description: string;
}

export interface NewInstanceRequest {
  name: string;
  description: string;
  schema_id: string;
}

export interface ApiError {
  error: string;
}
