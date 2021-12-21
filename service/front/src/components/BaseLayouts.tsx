import { Box, Container, Grid } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { ReactNode } from 'react'
import { Footer } from './Footer'
import { Header } from './Header'

type props = {
  main: ReactNode
}

export const BaseLayout: React.FC<props> = ({ main }) => {
  const classes = useStyles()
  return (
    <Box display="flex" flexDirection="column" minHeight="100vh">
      <Header />
      <Container className={classes.main}>
        <Grid container>{main}</Grid>
      </Container>
      <Footer />
    </Box>
  )
}

const useStyles = makeStyles(() => ({
  main: {
    flex: 1,
    overflowX: 'hidden',
  },
}))
