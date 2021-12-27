import { Grid } from '@mui/material'
import { NextPage } from 'next'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { PageProps } from '../../../types/page'
import { ProjectCreateForm } from '../../components/form/project/ProjectCreateForm'

const ProjectCreate: NextPage<PageProps> = (props) => {
  useRequireLogin()

  return (
    <Grid container mt={10}>
      <Grid xs={12} item>
        <ProjectCreateForm />
      </Grid>
    </Grid>
  )
}

export default ProjectCreate
