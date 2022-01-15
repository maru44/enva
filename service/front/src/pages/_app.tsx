import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { RecoilRoot, useRecoilState } from 'recoil'
import { Box, ThemeProvider } from '@mui/material'
import theme from '../theme/theme'
import { BaseLayout } from '../components/BaseLayouts'
import { SnackbarProvider } from 'notistack'
import { currentUserState } from '../../hooks/useCurrentUser'
import { useEffect } from 'react'
import { fetchCurrentUser } from '../../http/auth'
import Amplify, { Auth } from 'aws-amplify'
import awsConfig from '../../config/aws'
import { AmplifyProvider } from '@aws-amplify/ui-react'

Amplify.configure(awsConfig)
Auth.configure(awsConfig)

const AppInit = () => {
  const [, setCurrentUser] = useRecoilState(currentUserState)

  useEffect(() => {
    ;(async () => {
      try {
        const user = await fetchCurrentUser()
        setCurrentUser(user)
      } catch {
        setCurrentUser(null)
      }
    })()
  }, [setCurrentUser])

  return null
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <ThemeProvider theme={theme}>
        <SnackbarProvider maxSnack={2} autoHideDuration={5000}>
          <AmplifyProvider>
            <Box>
              <BaseLayout main={<Component {...pageProps} />} />
            </Box>
          </AmplifyProvider>
        </SnackbarProvider>
      </ThemeProvider>
      <AppInit />
    </RecoilRoot>
  )
}

export default MyApp
