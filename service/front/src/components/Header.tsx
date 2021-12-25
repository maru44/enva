import { Add, AddCircle } from '@material-ui/icons'
import { Grid, IconButton, Typography } from '@mui/material'
import Link from 'next/link'

export const Header: React.FC = () => {
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
        <Link href="/project/create" passHref>
          <IconButton>
            <AddCircle />
          </IconButton>
        </Link>
      </Grid>
    </Grid>
  )
}
