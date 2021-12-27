import { Box, Container, Grid } from '@mui/material'
import Head from 'next/head'
import { ReactNode } from 'react'
import { ThisUrl } from '../../config/env'
import { Footer } from './Footer'
import { Header } from './Header'

type props = {
  main: ReactNode
}

export const BaseLayout: React.FC<props> = ({ main }) => {
  return (
    <Box>
      <Head>
        <title>Envassador</title>
        <meta
          property="og:title"
          content="Envassador | Acceralate Local Development!"
        />
        <meta property="og:type" content="website" />
        {/* @TODO fix description */}
        <meta
          property="og:description"
          content="Envassador serves you an efficient system to share Environment variable. Envassador helps your local development."
        />
        <meta
          property="description"
          content="Envassador serves you an efficient system to share Environment variable. Envassador helps your local development."
        />
        <meta property="og:url" content={ThisUrl} />
        {/* @TODO: add image */}
        <meta property="og:image" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta
          name="twitter:title"
          content="Envassador | Acceralate Local Development!"
        />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Box
        display="flex"
        p={0}
        width="100%"
        flexDirection="column"
        minHeight="100vh"
        className="main"
      >
        <Header />
        <Container className="main">{main}</Container>
        <Footer />
      </Box>
    </Box>
  )
}
