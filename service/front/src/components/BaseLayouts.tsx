import { Box, Container, Grid } from '@mui/material'
import { makeStyles } from '@mui/styles'
import clsx from 'clsx'
import { ReactNode } from 'react'
import theme from '../theme/theme'
import { Footer } from './Footer'
import { Header } from './Header'

type props = {
  main: ReactNode
}

export const BaseLayout: React.FC<props> = ({ main }) => {
  const classes = useStyles(theme)
  return (
    <Box display="flex" flexDirection="column" minHeight="100vh">
      <Header />
      <Box className={clsx(classes.main)}>
        <Container>
          <Grid container>{main}</Grid>
        </Container>
      </Box>
      <Footer />
    </Box>
  )
}

const useStyles = makeStyles((theme) => ({
  main: {
    flex: 1,
    overflowX: 'hidden',
  },
}))
