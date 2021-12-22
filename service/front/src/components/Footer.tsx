import { Grid, Typography } from '@mui/material'
import Link from 'next/link'

export const Footer: React.FC = () => {
  return (
    <Grid
      container
      justifyContent="space-between"
      spacing={3}
      pt={1}
      pb={1}
      pl={2}
      mt={2}
    >
      <Grid item>
        <Typography>&copy; 2021 maru</Typography>
      </Grid>
    </Grid>
  )
}
