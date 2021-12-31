export type ProjectInput = {
  name: string
  slug: string
  description: string
  org_id?: string
}

// @TODO UNION with
export type Project = {
  id: string
  slug: string
  name: string
  description: string
  owner_type: string
  is_valid: boolean
  created_at: string
  updated_at: string
  deleted_at: string
}
