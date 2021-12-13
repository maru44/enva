import { ApiUrl } from '../config/env'
import { KvInput } from '../types/kv'
import { GetPath } from './fetcher'

export const fetchCreateKv = async (input: KvInput) => {
  return fetch(`${ApiUrl}${GetPath.KV_CREATE}`, {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: JSON.stringify(input),
  })
}
