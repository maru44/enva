import { ApiUrl } from '../config/env'
import { Project, ProjectInput } from '../types/project'

export const fetchProjectListByUser = async () => {
  return await fetch(`${ApiUrl}/project/list/user/`, {
    method: 'GET',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
  })
}

export const fetchProjectListByOrg = async (orgId: string) => {
  return await fetch(`${ApiUrl}/project/list/org/?orgId=${orgId}`, {
    method: 'GET',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
  })
}

export const fetchCreateProject = async (input: ProjectInput) => {
  return await fetch(`${ApiUrl}/project/create/`, {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: JSON.stringify(input),
  })
}
