import '../styles/globals.css'
import type { AppProps } from 'next/app'
import { RecoilRoot } from 'recoil'
import { Container } from '@mui/material'

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <Container maxWidth="lg">
        <Component {...pageProps} />
      </Container>
    </RecoilRoot>
  )
}

export default MyApp
