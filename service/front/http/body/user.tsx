import { CurrentUser } from '../../types/user'
import { errorResponseBody } from './error'

export type userResponseBody = {
  data: CurrentUser | null
} & errorResponseBody

export type userCRUDResponseBody = {
  data: string | null
} & errorResponseBody
