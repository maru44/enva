import { Box, Container, Grid } from '@mui/material'
import Head from 'next/head'
import { ReactNode } from 'react'
import { Footer } from './Footer'
import { Header } from './Header'

type props = {
  main: ReactNode
}

export const BaseLayout: React.FC<props> = ({ main }) => {
  return (
    <Box>
      <Head>
        <title>Envassador | Acceralate Development!</title>
        <meta
          property="og:title"
          content="Envassador | Acceralate Development!"
        />
        <meta property="og:type" content="website" />
        <meta
          property="og:description"
          content="Envassador serves you an efficient system to share environment variable. Envassador makes your development more comfortable."
        />
        <meta
          property="description"
          content="Envassador serves you an efficient system to share environment variable. Envassador makes your development more comfortable."
        />
        <meta property="og:url" content={process.env.NEXT_PUBLIC_FRONT_URL} />
        <meta
          property="og:image"
          content={`${process.env.NEXT_PUBLIC_FRONT_URL}/enva.png`}
        />
        <meta name="twitter:card" content="summary" />
        <meta
          name="twitter:title"
          content="Envassador | Acceralate Development!"
        />
        <link rel="icon" href="/enva.png" />
      </Head>
      <Box
        display="flex"
        p={0}
        width="100%"
        flexDirection="column"
        minHeight="100vh"
      >
        <Header />
        <Container className="main">{main}</Container>
        <Footer />
      </Box>
    </Box>
  )
}
