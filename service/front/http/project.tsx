import { ApiUrl } from '../config/env'
import { Project, ProjectInput } from '../types/project'

// @TODO delete
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

// @TODO delete
export const fetchCreateProject = async (input: ProjectInput) => {
  return await fetch(`${ApiUrl}/project/create`, {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: JSON.stringify(input),
  })
}
