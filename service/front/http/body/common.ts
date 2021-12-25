import { errorResponseBody } from './error'

export type deleteResponseBody = {
  data: string
} & errorResponseBody
