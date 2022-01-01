import { Org } from '../../types/org'
import { errorResponseBody } from './error'

export type OrgsResponseBody = {
  data: Org[]
} & errorResponseBody

export type OrgResponseBody = {
  data: Org
} & errorResponseBody

export type OrgCreateResponseBody = {
  data: string
} & errorResponseBody
