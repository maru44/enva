import {
  OrgInput,
  OrgInvitationInput,
  OrgMemberInput,
  OrgMemberUpdateInput,
} from '../types/org'
import { fetchBaseApi, GetPath } from './fetcher'

export const fetchCreateOrg = async (input: OrgInput) =>
  await fetchBaseApi(`${GetPath.ORG_CREATE}`, 'POST', input)

export const fetchOrgInvite = async (input: OrgInvitationInput) =>
  await fetchBaseApi(`${GetPath.ORG_INVITE}`, 'POST', input)

export const fetchAcceptInvitation = async (input: OrgMemberInput) =>
  await fetchBaseApi(`${GetPath.ORG_INVITATION_ACCEPT}`, 'POST', input)

export const fetchDenyInvitation = async (invId: string) =>
  await fetchBaseApi(`${GetPath.ORG_INVITATION_DENY}?id=${invId}`, 'GET')

export const fetchUpdateMemberUserType = async (input: OrgMemberUpdateInput) =>
  await fetchBaseApi(`${GetPath.ORG_MEMBER_UPDATE_TYPE}`, 'POST', input)

export const fetchDeleteMember = async (userId: string, orgId: string) =>
  await fetchBaseApi(
    `${GetPath.ORG_MEMBER_DELETE}?id=${userId}&orgId=${orgId}`,
    'DELETE'
  )
