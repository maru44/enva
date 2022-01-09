import { ArrowDropDown, ArrowDropUp } from '@material-ui/icons'
import { Box, Button, Icon, IconButton, Typography } from '@mui/material'
import { useSnackbar } from 'notistack'
import { useState } from 'react'
import { mutate } from 'swr'
import { OrgInviteResponseBody } from '../../../../http/body/org'
import { GetPath } from '../../../../http/fetcher'
import { fetchDeleteMember } from '../../../../http/org'
import { CurrentUser } from '../../../../types/user'

type props = {
  orgId: string
  user: CurrentUser
  onClose: () => void
}

export const MemberEliminateForm: React.FC<props> = ({
  orgId,
  user,
  onClose,
}) => {
  const [isOpen, setIsOpen] = useState<boolean>(false)
  const snack = useSnackbar()

  const onSubmit = async () => {
    const res = await fetchDeleteMember(user.id, orgId)
    const ret: OrgInviteResponseBody = await res.json()

    switch (res.status) {
      case 200:
        mutate(`${GetPath.ORG_MEMBERS_LIST}?id=${orgId}`)
        snack.enqueueSnackbar('Success to eliminate user', {
          variant: 'success',
        })
        onClose()
        break
      default:
        snack.enqueueSnackbar(ret.error, { variant: 'error' })
        break
    }
  }

  const openButton = (
    <Box>
      <Button
        type="button"
        onClick={() => {
          setIsOpen(!isOpen)
        }}
        fullWidth
        color="inherit"
      >
        <Icon>{isOpen ? <ArrowDropUp /> : <ArrowDropDown />}</Icon>
      </Button>
    </Box>
  )

  if (!isOpen) return openButton

  return (
    <Box>
      {openButton}
      <Box mt={2}>
        <Typography variant="subtitle1">
          Are you sure to eliminate {user.username}?
        </Typography>
      </Box>
      <Box textAlign="right" mt={2}>
        <Button variant="contained" type="button" onClick={onSubmit}>
          Confirm
        </Button>
      </Box>
    </Box>
  )
}
