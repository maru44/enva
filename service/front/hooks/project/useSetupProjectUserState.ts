import { atom } from 'recoil'
import { RecoilAtomKeys } from '../../config/recoil'
import { Project } from '../../types/project'

export const projectsState = atom({
  key: RecoilAtomKeys.PROJECT_LIST,
  default: {
    projects: [] as Project[],
  },
})

// export const useSetupProjectUserState = () => {
//     const
// }
