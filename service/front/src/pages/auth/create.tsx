import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { useSnackbar } from 'notistack'
import { useEffect } from 'react'
import { userCreateResponseBody } from '../../../http/body/user'
import { fetchBaseApi, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'

const UserCreate: NextPage<PageProps> = (props) => {
  const router = useRouter()
  const snack = useSnackbar()

  useEffect(() => {
    ;(async () => {
      // insert user to db
      const res = await fetchBaseApi(GetPath.USER_CREATE, 'GET')
      const ret: userCreateResponseBody = await res.json()
      if (res.status == 200) {
        snack.enqueueSnackbar('signed in!', { variant: 'success' })
        router.push('/project')
      } else {
        snack.enqueueSnackbar(ret.error, { variant: 'error' })
        router.push('/')
      }
    })()
  }, [router])

  return <div></div>
}

export default UserCreate
