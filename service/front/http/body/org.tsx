import { Org, OrgInvitation } from '../../types/org'
import { UserType } from '../../types/user'
import { errorResponseBody } from './error'

export type OrgsResponseBody = {
  data: Org[]
} & errorResponseBody

export type OrgResponseBody = {
  data: {
    org: Org
    current_user_type: UserType
  }
} & errorResponseBody

export type OrgCreateResponseBody = {
  data: string
} & errorResponseBody

// invite / create member / deny
export type OrgInviteResponseBody = {
  data: string
} & errorResponseBody

export type OrgInvitationDetailBody = {
  data: OrgInvitation
} & errorResponseBody
