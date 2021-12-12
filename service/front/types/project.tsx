export type ProjectInput = {
  name: string
  slug: string
  description: string
  org_id?: string
}

// @TODO UNION
export type Project = {
  id: string
  slug: string
  name: string
  owner_type: string
  is_valid: boolean
  is_deleted: boolean
  created_at: string
  updated_at: string
}
