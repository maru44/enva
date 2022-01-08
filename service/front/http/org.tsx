import { OrgInput, OrgInvitationInput } from '../types/org'
import { fetchBaseApi, GetPath } from './fetcher'

export const fetchCreateOrg = async (input: OrgInput) =>
  await fetchBaseApi(`${GetPath.ORG_CREATE}`, 'POST', input)

export const fetchOrgInvite = async (input: OrgInvitationInput) =>
  await fetchBaseApi(`${GetPath.ORG_INVITE}`, 'POST', input)
