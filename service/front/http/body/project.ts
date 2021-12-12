import { Project } from '../../types/project'
import { errorResponseBody } from './error'

export type projectsResponseBody =
  | {
      data: Project[]
    } & errorResponseBody

export type projectResponseBody =
  | {
      data: Project
    } & errorResponseBody
