import { Grid } from '@mui/material'
import { NextPage } from 'next'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { PageProps } from '../../../types/page'
import { OrgCreateForm } from '../../components/form/org/OrgCreateForm'

const OrgCreate: NextPage<PageProps> = (props) => {
  useRequireLogin()

  return (
    <Grid container mt={10}>
      <Grid xs={12} item>
        <OrgCreateForm />
      </Grid>
    </Grid>
  )
}

export default OrgCreate
