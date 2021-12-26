import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { RecoilRoot, useRecoilState } from 'recoil'
import { Box, Theme, ThemeProvider } from '@mui/material'
import theme from '../theme/theme'
import { BaseLayout } from '../components/BaseLayouts'
import { SnackbarProvider } from 'notistack'
import { currentUserState } from '../../hooks/useCurrentUser'
import { useEffect } from 'react'
import { fetchCurrentUser } from '../../http/auth'

declare module '@mui/styles/defaultTheme' {
  interface DefaultTheme extends Theme {}
}

const AppInit = () => {
  const [, setCurrentUser] = useRecoilState(currentUserState)

  useEffect(() => {
    ;(async () => {
      try {
        const user = await fetchCurrentUser()
        console.log(user)
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
          <Box>
            <BaseLayout main={<Component {...pageProps} />} />
          </Box>
        </SnackbarProvider>
      </ThemeProvider>
      <AppInit />
    </RecoilRoot>
  )
}

export default MyApp
