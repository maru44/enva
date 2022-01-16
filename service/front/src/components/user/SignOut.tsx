import { Box, Button, Grid, Paper, Typography } from '@mui/material'
import { useSnackbar } from 'notistack'
import { useRecoilState } from 'recoil'
import { ThisUrl } from '../../../config/env'
import { currentUserState } from '../../../hooks/useCurrentUser'

export const SignOut: React.FC = () => {
  const snack = useSnackbar()
  const [, setCurrentUser] = useRecoilState(currentUserState)
  const onSignOut = async () => {
    try {
      const res = await fetch(`${ThisUrl}/api/auth/signout`, {
        method: 'GET',
        credentials: 'include',
      })
      if (res.status === 200) {
        snack.enqueueSnackbar('signed out', { variant: 'info' })
        setCurrentUser(null)
        return
      }
      snack.enqueueSnackbar('Internal Server Error', { variant: 'error' })
    } catch {
      snack.enqueueSnackbar('Internal Server Error', { variant: 'error' })
    }
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
