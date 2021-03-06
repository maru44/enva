import {
  Box,
  Button,
  Dialog,
  Grid,
  MenuItem,
  Select,
  TextField,
  Typography,
} from '@mui/material'
import { useSnackbar } from 'notistack'
import { FormEvent, useState } from 'react'
import { mutate } from 'swr'
import { OrgInviteResponseBody } from '../../../../http/body/org'
import { GetPath } from '../../../../http/fetcher'
import { fetchOrgInvite } from '../../../../http/org'
import { OrgInvitationInput } from '../../../../types/org'
import { UserType, UserTypesAll } from '../../../../types/user'

type props = {
  orgId: string
  orgName: string
  isOpen: boolean
  onClose: () => void
}

export const InviteFormModal: React.FC<props> = ({
  orgName,
  orgId,
  isOpen,
  onClose,
}) => {
  const [isPosting, setIsPosting] = useState<boolean>(false)
  const snack = useSnackbar()

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    try {
      e.preventDefault()
      setIsPosting(true)
      const email = e.currentTarget.email.value
      const userType = e.currentTarget.user_type.value

      const input: OrgInvitationInput = {
        org_id: orgId,
        org_name: orgName,
        email: email,
        user_type: userType,
      }

      const res = await fetchOrgInvite(input)
      const ret: OrgInviteResponseBody = await res.json()

      switch (res.status) {
        case 200:
          snack.enqueueSnackbar(`success to invite ${email}`, {
            variant: 'success',
          })
          mutate(`${GetPath.ORG_INVITATION_LIST}?id=${orgId}`)
          onClose()
          break
        default:
          snack.enqueueSnackbar(ret.error, {
            variant: 'error',
          })
          break
      }
      setIsPosting(false)
    } catch (e) {
      snack.enqueueSnackbar('Internal Server Error', { variant: 'error' })
    }
  }

  return (
    <Dialog onClose={onClose} open={isOpen}>
      <Grid container className="dialogContainer" p={3}>
        <Grid sm={2} item />
        <Grid
          component="form"
          xs={8}
          item
          onSubmit={(e: FormEvent<HTMLFormElement>) => onSubmit(e)}
        >
          <Box mt={2}>
            <Typography variant="h5">Inviting to {orgName}</Typography>
          </Box>
          <Box mt={4}>
            <TextField
              required
              fullWidth
              label="email"
              type="email"
              name="email"
            ></TextField>
          </Box>
          <Box mt={4}>
            <Select
              label="user type"
              name="user_type"
              fullWidth
              defaultValue={UserType.USER}
            >
              {UserTypesAll.map((ut, i) => (
                <MenuItem key={i} value={ut}>
                  {ut}
                </MenuItem>
              ))}
            </Select>
          </Box>
          <Box mt={4} display="flex" justifyContent="space-between">
            <Box>
              <Button
                onClick={onClose}
                type="button"
                variant="contained"
                className="subButton"
                disabled={isPosting}
              >
                Close
              </Button>
            </Box>
            <Box>
              <Button disabled={isPosting} type="submit" variant="contained">
                Invite
              </Button>
            </Box>
          </Box>
        </Grid>
      </Grid>
    </Dialog>
  )
}
