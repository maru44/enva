import { Box } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import { useSnackbar } from 'notistack'
import { useRecoilState } from 'recoil'
import { currentUserState } from '../../../hooks/useCurrentUser'
import { PageProps } from '../../../types/page'

const SignOut: NextPage<PageProps> = (props) => {
  const snack = useSnackbar()
  const router = useRouter()
  const [, setCurrentUser] = useRecoilState(currentUserState)

  if (router.isReady) {
    ;(async () => {
      try {
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_FRONT_URL}/api/auth/signout`
        )
        switch (res.status) {
          case 200:
            setCurrentUser(null)
            snack.enqueueSnackbar('Signed out', { variant: 'info' })
            break
          default:
            snack.enqueueSnackbar('Failed to sign out', { variant: 'error' })
            break
        }
      } catch {
        snack.enqueueSnackbar('Failed to sign out', { variant: 'error' })
      }
    })()
    router.push('/')
  }

  return <Box></Box>
}

export default SignOut
