import {
  Box,
  Card,
  Dialog,
  Grid,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
} from '@mui/material'
import useSWR from 'swr'
import { OrgInvitationListResponseBody } from '../../../../http/body/org'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'
import { InvitationStatusBadge } from './InvitationStatusBadge'
import styles from '../../../styles/org.module.css'

type props = {
  orgId: string
  isOpen: boolean
  onClose: () => void
}

export const InvitationHistoryModal: React.FC<props> = ({
  orgId,
  isOpen,
  onClose,
}) => {
  const { data, error } = useSWR<
    OrgInvitationListResponseBody,
    ErrorConstructor
  >(`${GetPath.ORG_INVITATION_LIST}?id=${orgId}`, fetcherGetFromApiUrl)

  if (!data?.data) return <></>

  return (
    <Dialog open={isOpen} onClose={onClose} maxWidth="lg">
      <Grid container className="dialogContainerLarge" pt={3} pb={3}>
        <Grid xs={1} item />
        <Grid xs={10} item p={0}>
          <Box>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell width="52%">
                    <b>Target</b>
                  </TableCell>
                  <TableCell width="16%">
                    <b>Type</b>
                  </TableCell>
                  <TableCell width="16%">
                    <b>Status</b>
                  </TableCell>
                  {/* <TableCell width="16%">Invited @</TableCell> */}
                </TableRow>
              </TableHead>
              <TableBody>
                {data.data.map((inv, i) => (
                  <TableRow key={i}>
                    <TableCell>
                      <Typography className="breakAll">
                        {inv.user.username !== ''
                          ? inv.user.username
                          : inv.user.email}
                      </Typography>
                    </TableCell>
                    <TableCell>{inv.user_type}</TableCell>
                    <TableCell>
                      <InvitationStatusBadge status={inv.status} />
                    </TableCell>
                    {/* <TableCell>
                      <Typography className="breakAll">
                        {inv.created_at}
                      </Typography>
                    </TableCell> */}
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </Box>
        </Grid>
      </Grid>
    </Dialog>
  )
}
