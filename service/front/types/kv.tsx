export type Kv = {
  id: string
  kv_key: string
  kv_value: string
}

export type KvInput = {
  project_id: string
  input: {
    kv_key: string
    kv_value: string
  }
}
