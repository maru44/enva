import { KvInput } from '../types/kv'
import { fetchBaseApi, GetPath } from './fetcher'

export const fetchCreateKv = async (input: KvInput) => {
  return fetchBaseApi(GetPath.KV_CREATE, 'POST', input)
}

export const fetchUpdateKv = async (input: KvInput) => {
  return fetchBaseApi(GetPath.KV_UPDATE, 'PUT', input)
}

export const fetchDeleteKv = async (kvId: string, projectId: string) => {
  return fetchBaseApi(`${GetPath.KV_DELETE}?kvId=${kvId}&projectId=${projectId}`, 'DELETE')
}