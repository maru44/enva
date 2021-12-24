import { Box } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { NextPage } from 'next'
import { PageProps } from '../../../types/page'
import { ProjectCreateForm } from '../../components/form/project/ProjectCreateForm'

const ProjectCreate: NextPage<PageProps> = (props) => {
  return (
    <Box m={2}>
      <Box>
        <ProjectCreateForm />
      </Box>
    </Box>
  )
}

const useStyle = makeStyles(() => ({
  root: {
    // padding: theme.spacing(1),
  },
}))

export default ProjectCreate
