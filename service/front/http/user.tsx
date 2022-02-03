import { fetchBaseApi, GetPath } from './fetcher'

export const fetchWithdraw = async () =>
  await fetchBaseApi(`${GetPath.USER_WITHDRAW}`, 'GET')
