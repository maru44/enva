import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { RecoilRoot } from 'recoil'
import { Box, Container, Theme, ThemeProvider } from '@mui/material'
import theme from '../theme/theme'
import { BaseLayout } from '../components/BaseLayouts'

declare module '@mui/styles/defaultTheme' {
  interface DefaultTheme extends Theme {}
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <ThemeProvider theme={theme}>
        <Container maxWidth={false}>
          {/* <Component {...pageProps} /> */}
          <BaseLayout main={<Component {...pageProps} />} />
        </Container>
      </ThemeProvider>
    </RecoilRoot>
  )
}

export default MyApp
