import {
  Box,
  Dialog,
  Grid,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableRow,
  Typography,
} from '@mui/material'
import useSWR from 'swr'
import { OrgMemberTypeResponseBody } from '../../../../http/body/org'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'
import { CurrentUser, UserType } from '../../../../types/user'
import { MemberEliminateForm } from './MemberEliminateForm'
import { UpdateUserTypeForm } from './UpdateUserTypeForm'

type props = {
  user?: CurrentUser
  orgId: string
  defaultType?: UserType
  currentUserType: UserType
  onClose: () => void
}

export const MemberDetailModal: React.FC<props> = ({
  user,
  orgId,
  defaultType,
  currentUserType,
  onClose,
}) => {
  const { data, error } = useSWR<OrgMemberTypeResponseBody, ErrorConstructor>(
    `${GetPath.ORG_MEMBER_TYPE}?id=${user?.id}&orgId=${orgId}`,
    fetcherGetFromApiUrl
  )

  if (!user) return <></>

  const canEdit =
    currentUserType === UserType.OWNER ||
    (currentUserType === UserType.ADMIN && defaultType !== UserType.OWNER)

  return (
    <Dialog onClose={onClose} open={true}>
      <Grid container className="dialogContainer" p={3}>
        <Grid sm={2} item />
        <Grid item xs={8}>
          <Grid item xs={12}>
            <Box>
              <Typography variant="h5">Member</Typography>
            </Box>
            <Box mt={2}>
              <Grid component={Paper} item>
                <Table>
                  <TableBody>
                    <TableRow>
                      <TableCell>Username</TableCell>
                      <TableCell>{user.username}</TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell>Email</TableCell>
                      <TableCell>{user.email}</TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell>User Type</TableCell>
                      <TableCell>{data?.data ?? defaultType}</TableCell>
                    </TableRow>
                  </TableBody>
                </Table>
              </Grid>
            </Box>
            {canEdit && (
              <Box mt={4}>
                <UpdateUserTypeForm
                  orgId={orgId}
                  defaultType={data?.data ?? defaultType!}
                  user={user!}
                />
                <Box mt={2}>
                  <Box>
                    <Typography variant="h5">Eliminate User</Typography>
                  </Box>
                  <Box mt={2}>
                    <MemberEliminateForm
                      user={user}
                      orgId={orgId}
                      onClose={onClose}
                    />
                  </Box>
                </Box>
              </Box>
            )}
          </Grid>
        </Grid>
      </Grid>
    </Dialog>
  )
}
