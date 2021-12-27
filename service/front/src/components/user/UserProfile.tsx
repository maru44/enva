import {
  Box,
  Grid,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableRow,
  Typography,
} from '@mui/material'
import { CurrentUser } from '../../../types/user'

type props = {
  currentUser: CurrentUser
}

export const UserProfile: React.FC<props> = ({ currentUser }) => {
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
        <Typography variant="h5">Profile</Typography>
        <Box mt={2}>
          <TableContainer>
            <Table>
              <TableBody>
                <TableRow>
                  <TableCell>Username</TableCell>
                  <TableCell>{currentUser.username}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Email</TableCell>
                  <TableCell>{currentUser.email}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Has cli password</TableCell>
                  <TableCell>
                    {currentUser.has_cli_password ? 'true' : 'false'}
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
      </Grid>
    </Grid>
  )
}
