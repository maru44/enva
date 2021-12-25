import { CurrentUser } from '../../types/user'
import { errorResponseBody } from './error'

export type userResponseBody = {
  data: CurrentUser
} & errorResponseBody
