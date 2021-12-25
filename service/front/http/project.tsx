import { ApiUrl } from '../config/env'
import { Project, ProjectInput } from '../types/project'
import { fetchBaseApi, GetPath } from './fetcher'

export const fetchProjectListByUser = async () => {
  return await fetch(`${ApiUrl}/project/list/user`, {
    method: 'GET',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
  })
}

export const fetchCreateProject = async (input: ProjectInput) =>
  await fetchBaseApi(GetPath.PROJECT_CREATE, 'POST', input)

// export const fetchUpdateProject = async (input: ProjectInput) => {}

export const fetchDeleteProject = async (projectId: string) =>
  await fetchBaseApi(
    `${GetPath.PROJECT_DELETE}/project/delete/?projectId=${projectId}`,
    'DELETE'
  )
