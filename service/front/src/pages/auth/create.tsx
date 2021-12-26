import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { fetchBaseApi, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'

const UserCreate: NextPage<PageProps> = (props) => {
  const router = useRouter()

  useEffect(() => {
    ;(async () => {
      // insert user to db
      const res = await fetchBaseApi(GetPath.USER_CREATE, 'GET')
      router.push('/project')
    })()
  }, [router])

  return <div></div>
}

export default UserCreate
