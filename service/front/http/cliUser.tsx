import { CliUserValidateInput } from '../types/cliUser'
import { fetchBaseApi, GetPath } from './fetcher'

export const fetchUpdateCliUser = async () =>
  await fetchBaseApi(GetPath.CLI_USER_UPDATE, 'GET')

export const fetchValidateCliUser = async (input: CliUserValidateInput) =>
  await fetchBaseApi(GetPath.CLI_USER, 'POST', input)
