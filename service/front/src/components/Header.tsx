import { AccountCircle, Add, AddCircle } from '@material-ui/icons'
import { Box, Button, Grid, IconButton, Typography } from '@mui/material'
import Link from 'next/link'
import { loginUrl } from '../../config/aws'
import { useCurrentUser } from '../../hooks/useCurrentUser'

export const Header: React.FC = () => {
  const { isAuthChecking, currentUser } = useCurrentUser()

  return (
    <Grid
      className="header"
      container
      justifyContent="space-between"
      spacing={3}
      pt={1}
      pb={1}
    >
      <Grid item>
        <Typography className="title" variant="h4" pl={1}>
          <Link href={currentUser ? '/project' : '/'} passHref>
            <a>Envassador</a>
          </Link>
        </Typography>
      </Grid>
      <Grid item>
        {currentUser && (
          <Box display="flex" flexDirection="row">
            <Box mr={1}>
              <Link href="/user" passHref>
                <IconButton color="primary">
                  <AccountCircle />
                </IconButton>
              </Link>
            </Box>
            <Box mr={1}>
              <Link href="/project/create" passHref>
                <IconButton color="primary">
                  <AddCircle />
                </IconButton>
              </Link>
            </Box>
          </Box>
        )}
        {!isAuthChecking && !currentUser && (
          <Box mr={2}>
            <Link href={loginUrl} passHref>
              <Button>Sign in</Button>
            </Link>
          </Box>
        )}
      </Grid>
    </Grid>
  )
}
