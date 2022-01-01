import { BaseInformation } from './information'
import { CurrentUser } from './user'

export type Org = {
  is_valid: boolean
  created_by: CurrentUser
  created_at: string
  updated_at: string

  user_count: number
} & BaseInformation

export type OrgInput = {
  slug: string
  name: string
  description?: string
}
