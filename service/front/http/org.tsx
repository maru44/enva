import { OrgInput } from '../types/org'
import { fetchBaseApi, GetPath } from './fetcher'

export const fetchCreateOrg = async (input: OrgInput) =>
  await fetchBaseApi(`${GetPath.ORG_CREATE}`, 'POST', input)
