import { Box, Grid, Typography } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { NextPage } from 'next'
import { PageProps } from '../../../types/page'
import { ProjectCreateForm } from '../../components/form/project/ProjectCreateForm'

const ProjectCreate: NextPage<PageProps> = (props) => {
  return (
    <Grid container mt={2}>
      <Grid xs={12} item>
        <ProjectCreateForm />
      </Grid>
    </Grid>
  )
}

const useStyle = makeStyles(() => ({
  root: {
    // padding: theme.spacing(1),
  },
}))

export default ProjectCreate
