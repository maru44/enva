import { Box, Button, Grid, Paper, Typography } from '@mui/material'
import { useRouter } from 'next/router'
import { logoutUrl } from '../../../config/aws'

export const SignOut: React.FC = () => {
  const router = useRouter()
  const onSignOut = () => {
    router.push(logoutUrl)
  }

  return (
    <Grid container>
      <Grid item xs={0} sm={2} md={3} />
      <Grid
        item
        xs={12}
        sm={8}
        md={6}
        component={Paper}
        pt={2}
        pb={2}
        pl={1}
        pr={1}
        variant="outlined"
      >
        <Typography variant="h5">Sign Out</Typography>
        <Box mt={2}>
          <Typography>Are you sure to sign out?</Typography>
          <Box mt={2} textAlign="right">
            <Button onClick={onSignOut} type="button" variant="contained">
              Sign out
            </Button>
          </Box>
        </Box>
      </Grid>
    </Grid>
  )
}
