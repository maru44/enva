import {
  AccountCircle,
  Add,
  AddCircle,
  Computer,
  ComputerOutlined,
  ComputerRounded,
} from '@material-ui/icons'
import {
  Box,
  Button,
  Grid,
  IconButton,
  Tooltip,
  Typography,
} from '@mui/material'
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
              <Link href="/cli" passHref>
                <Tooltip arrow title="download cli">
                  <IconButton color="primary">
                    <ComputerOutlined />
                  </IconButton>
                </Tooltip>
              </Link>
            </Box>
            <Box mr={1}>
              <Link href="/user" passHref>
                <Tooltip arrow title="user info">
                  <IconButton color="primary">
                    <AccountCircle />
                  </IconButton>
                </Tooltip>
              </Link>
            </Box>
            <Box mr={1}>
              <Link href="/project/create" passHref>
                <Tooltip arrow title="add project">
                  <IconButton color="primary">
                    <AddCircle />
                  </IconButton>
                </Tooltip>
              </Link>
            </Box>
          </Box>
        )}
        {!isAuthChecking && !currentUser && (
          <Box display="flex" flexDirection="row">
            <Box mr={1}>
              <Link href="/cli" passHref>
                <Tooltip arrow title="download cli">
                  <IconButton color="primary">
                    <ComputerOutlined />
                  </IconButton>
                </Tooltip>
              </Link>
            </Box>
            <Box mr={2}>
              <Link href={loginUrl} passHref>
                <Button>Sign in</Button>
              </Link>
            </Box>
          </Box>
        )}
      </Grid>
    </Grid>
  )
}
