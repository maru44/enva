import { AccountCircle, Add, AddCircle } from '@material-ui/icons'
import { Box, Grid, IconButton, Typography } from '@mui/material'
import Link from 'next/link'
import { useCurrentUser } from '../../hooks/useCurrentUser'

export const Header: React.FC = () => {
  const { currentUser } = useCurrentUser()

  return (
    <Grid container justifyContent="space-between" spacing={3} pt={1} pb={1}>
      <Grid item>
        <Typography variant="h4" pl={1}>
          <Link href="/project" passHref>
            <a>Envassador</a>
          </Link>
        </Typography>
      </Grid>
      <Grid item>
        <Box display="flex" flexDirection="row">
          <Box>
            {/* {currentUser ? } */}
            <Link href="/auth/profile" passHref>
              <IconButton>
                <AccountCircle />
              </IconButton>
            </Link>
          </Box>
          <Box>
            <Link href="/project/create" passHref>
              <IconButton>
                <AddCircle />
              </IconButton>
            </Link>
          </Box>
        </Box>
      </Grid>
    </Grid>
  )
}
