export type Kv = {
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
