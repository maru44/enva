import { BaseInformation } from './information'

export type ProjectInput = {
  name: string
  slug: string
  description: string
  org_id?: string
}

export type Project = {
  owner_type: string
  is_valid: boolean
  created_at: string
  updated_at: string
  deleted_at: string
} & BaseInformation
