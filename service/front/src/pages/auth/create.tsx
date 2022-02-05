import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { useSnackbar } from 'notistack'
import { useEffect } from 'react'
import { useRecoilState } from 'recoil'
import { currentUserState } from '../../../hooks/useCurrentUser'
import { userResponseBody } from '../../../http/body/user'
import { fetchBaseApi, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'

const UserCreate: NextPage<PageProps> = (props) => {
  const router = useRouter()
  const snack = useSnackbar()
  const [, setCurrentUser] = useRecoilState(currentUserState)

  useEffect(() => {
    ;(async () => {
      // insert user to db
      const res = await fetchBaseApi(GetPath.USER_CREATE, 'GET')
      const ret: userResponseBody = await res.json()
      if (res.status === 200) {
        snack.enqueueSnackbar('signed in!', { variant: 'success' })
        setCurrentUser(ret.data)
        router.push('/project')
      } else {
        snack.enqueueSnackbar(ret.error, { variant: 'error' })
        setCurrentUser(null)
        router.push('/')
      }
    })()
  }, [router, snack, setCurrentUser])

  return <div></div>
}

export default UserCreate
