import { Kv } from '../../types/kv'
import { errorResponseBody } from './error'

export type kvsResponseBody = {
  data: Kv[]
} & errorResponseBody

export type kvCreateResponseBody = {
  data: string
} & errorResponseBody

export type kvDeleteResponseBody = {
  data: string
} & errorResponseBody
