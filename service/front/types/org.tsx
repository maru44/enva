import { BaseInformation } from './information'
import { CurrentUser, UserType } from './user'

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

export type OrgInvitationInput = {
  org_id: string
  org_name: string
  email: string
  user_type: UserType
}

export type OrgInvitation = {
  id: string
  user_type: UserType
  status: InvitationStatus
  created_at: string
  updated_at: string

  org: Org
  user: CurrentUser
  invitor: CurrentUser
}

export type OrgMemberInput = {
  org_id: string
  user_id: string
  user_type: UserType
  org_invitation_id: string
}

export const InvitationStatus = {
  NEW: 'new',
  ACCEPTED: 'accepted',
  CLOSED: 'closed',
  DENIED: 'denied',
}
export type InvitationStatus =
  typeof InvitationStatus[keyof typeof InvitationStatus]
