import { Org } from '../../types/org'
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

export type OrgInviteResponseBody = {
  data: string
} & errorResponseBody
