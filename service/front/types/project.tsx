import { BaseInformation } from './information'
import { Org } from './org'
import { CurrentUser } from './user'

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
  user?: CurrentUser
  org?: Org
} & BaseInformation
