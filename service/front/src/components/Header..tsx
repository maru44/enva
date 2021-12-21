import { Grid, Typography } from '@mui/material'
import Link from 'next/link'

export const Header: React.FC = () => {
  return (
    <Grid
      container
      justifyContent="space-between"
      spacing={3}
      pt={1}
      pb={1}
      pl={2}
    >
      <Grid item>
        <Typography variant="h4">
          <Link href="/project" passHref>
            <a>Envassador</a>
          </Link>
        </Typography>
      </Grid>
    </Grid>
  )
}
