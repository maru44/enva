import { ApiUrl } from '../config/env'
import { KvInput } from '../types/kv'
import { fetchBaseApi, GetPath } from './fetcher'

export const fetchCreateKv = async (input: KvInput) => {
  //   return fetch(`${ApiUrl}${GetPath.KV_CREATE}`, {
  //     method: 'POST',
  //     mode: 'cors',
  //     credentials: 'include',
  //     headers: {
  //       'Content-Type': 'application/json; charset=utf-8',
  //     },
  //     body: JSON.stringify(input),
  //   })
  return fetchBaseApi(GetPath.KV_CREATE, 'POST', input)
}

export const fetchUpdateKv = async (input: KvInput) => {
  return fetch(`${ApiUrl}${GetPath.KV_UPDATE}`, {
    method: 'PUT',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: JSON.stringify(input),
  })
}
