import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { RecoilRoot } from 'recoil'
import { Container, Theme, ThemeProvider } from '@mui/material'
import theme from '../theme/theme'
import { BaseLayout } from '../components/BaseLayouts'
import { SnackbarProvider } from 'notistack'

declare module '@mui/styles/defaultTheme' {
  interface DefaultTheme extends Theme {}
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <ThemeProvider theme={theme}>
        <SnackbarProvider maxSnack={2} autoHideDuration={5000}>
          <Container maxWidth={false}>
            <BaseLayout main={<Component {...pageProps} />} />
          </Container>
        </SnackbarProvider>
      </ThemeProvider>
    </RecoilRoot>
  )
}

export default MyApp
