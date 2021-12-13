import { Kv } from '../../types/kv'
import { errorResponseBody } from './error'

export type kvsResponseBody = {
  data: Kv[]
} & errorResponseBody

export type kvCreateResponseBody = {
  data: {
    env_key: string
    env_value: string
  }
} & errorResponseBody
