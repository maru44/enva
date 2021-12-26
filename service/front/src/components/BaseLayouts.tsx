import { Box, Container, Grid } from '@mui/material'
import { ReactNode } from 'react'
import { Footer } from './Footer'
import { Header } from './Header'

type props = {
  main: ReactNode
}

export const BaseLayout: React.FC<props> = ({ main }) => {
  return (
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
  )
}
