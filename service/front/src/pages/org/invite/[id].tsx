import { Box, Button, Grid, Paper, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { useRequireLogin } from '../../../../hooks/useRequireLogin'
import {
  errorResponseBody,
  internalServerErrorInFront,
} from '../../../../http/body/error'
import {
  OrgInvitationDetailBody,
  OrgInviteResponseBody,
} from '../../../../http/body/org'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'
import {
  fetchAcceptInvitation,
  fetchDenyInvitation,
} from '../../../../http/org'
import { OrgMemberInput } from '../../../../types/org'
import { ErrorComponent } from '../../../components/error/ErrorComponent'

const OrgInvitationDetailPage: NextPage = () => {
  useRequireLogin()

  const router = useRouter()
  const id = router.query.id

  const { data, error } = useSWR<OrgInvitationDetailBody, ErrorConstructor>(
    `${GetPath.ORG_INVITATION_DETAIL}?id=${id}`,
    fetcherGetFromApiUrl
  )

  const deny = async () => {
    const res = await fetchDenyInvitation(id as string)
    const ret: OrgInviteResponseBody = await res.json()

    switch (res.status) {
      case 200:
        // @TODO snack
        router.push('/project')
        break
      default:
        // @TODO snack
        break
    }
  }

  const accept = async () => {
    const input: OrgMemberInput = {
      org_id: data?.data.org.id!,
      user_id: data?.data.user.id!,
      user_type: data?.data.user_type!,
      org_invitation_id: id as string,
    }

    const res = await fetchAcceptInvitation(input)
    const ret: OrgInviteResponseBody = await res.json()

    switch (res.status) {
      case 200:
        //   @TODO snack
        router.push(`/org/${data?.data.org.slug}`)
        break
      default:
        // @TODO snack
        break
    }
  }

  if (error) return <ErrorComponent />
  if (data?.error) return <ErrorComponent errBody={data} />

  if (data?.data)
    return (
      <Box>
        <Grid container mt={6}>
          <Grid item xs={0} sm={1} md={3} />
          <Grid component={Paper} item xs={12} sm={10} md={6} p={1}>
            <Box>
              <Typography variant="h5">
                Invitation to {data.data.org.name}
              </Typography>
            </Box>
            {data.data.org.description && (
              <Box mt={4}>
                <Typography>{data.data.org.description}</Typography>
              </Box>
            )}
            <Box mt={4}>
              <Typography>
                You are invited by {data.data.invitor.username}.
              </Typography>
            </Box>
            <Box
              mt={4}
              display="flex"
              flexDirection="row"
              justifyContent="space-evenly"
            >
              <Box>
                <Button
                  type="button"
                  variant="contained"
                  color="warning"
                  onClick={deny}
                >
                  Deny
                </Button>
              </Box>
              <Box>
                <Button type="button" variant="contained" onClick={accept}>
                  Accept
                </Button>
              </Box>
            </Box>
          </Grid>
        </Grid>
      </Box>
    )

  return <Box></Box>
}

export default OrgInvitationDetailPage
