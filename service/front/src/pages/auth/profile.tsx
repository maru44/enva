import { Box, Grid, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useCurrentUser } from '../../../hooks/useCurrentUser'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { PageProps } from '../../../types/page'

const AuthProfile: NextPage<PageProps> = (props) => {
  useRequireLogin()

  const { currentUser } = useCurrentUser()

  return (
    <Box mt={2} width="100%">
      <Typography variant="h5">Profile</Typography>
      <Grid mt={1}>
        <Box>
          {currentUser?.username}
          {currentUser?.email}
        </Box>
      </Grid>
    </Box>
  )
}

export default AuthProfile
