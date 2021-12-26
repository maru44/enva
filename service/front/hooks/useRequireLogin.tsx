import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { loginUrl } from '../config/aws'
import { useCurrentUser } from './useCurrentUser'

export const useRequireLogin = () => {
  const { isAuthChecking, currentUser } = useCurrentUser()
  const router = useRouter()

  useEffect(() => {
    if (isAuthChecking) return
    if (!currentUser) router.push(loginUrl)
  }, [isAuthChecking, currentUser, router])
}
