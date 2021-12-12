import { selector, useRecoilValue } from 'recoil'
import { RecoilSelectorKeys } from '../../config/recoil'
import { projectsResponseBody } from '../../http/body/project'
import { fetcher } from '../../http/fetcher'
import { fetchProjectListByUser } from '../../http/project'
import { projectsState } from './useSetupProjectUserState'

export const projectsUserState = selector({
  key: RecoilSelectorKeys.PROJECT_LIST_USER,
  get: async ({ get }) => {
    const { projects } = get(projectsState)
    const res = await fetcher(fetchProjectListByUser())
    if (res.status === 200) {
      const ret: projectsResponseBody = await res.json()
      const newProjects = ret.data

      if (newProjects !== projects) return { newProjects }
    }
    return { projects }
  },
})

export const useProjectUser = () => {
  return useRecoilValue(projectsUserState)
}
