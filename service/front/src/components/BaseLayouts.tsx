import { Box, Container, Grid } from '@mui/material'
import { ReactNode } from 'react'
import { Header } from './Header.'

type props = {
  main: ReactNode
}

export const BaseLayout: React.FC<props> = ({ main }) => {
  return (
    <Box>
      <Header />
      <Container>
        <Grid container>{main}</Grid>
      </Container>
    </Box>
  )
}
