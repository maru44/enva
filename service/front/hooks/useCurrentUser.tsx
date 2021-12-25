import { atom, useRecoilValue } from 'recoil'
import { CurrentUser } from '../types/user'

// undefined >> not confirmed yet
// null >> not signed in
export const currentUserState = atom<undefined | null | CurrentUser>({
  key: 'CurrentUser',
  default: undefined,
})

export const useCurrentUser = () => {
  const currentUser = useRecoilValue(currentUserState)
  const isAuthChecking = currentUser === undefined

  return {
    currentUser,
    isAuthChecking,
  }
}
