import {
  Box,
  Button,
  MenuItem,
  Select,
  SelectChangeEvent,
  Typography,
} from '@mui/material'
import { useSnackbar } from 'notistack'
import React, { useState } from 'react'
import { mutate } from 'swr'
import { OrgInviteResponseBody } from '../../../../http/body/org'
import { GetPath } from '../../../../http/fetcher'
import { fetchUpdateMemberUserType } from '../../../../http/org'
import { OrgMemberUpdateInput } from '../../../../types/org'
import { CurrentUser, UserType, UserTypesAll } from '../../../../types/user'

type props = {
  defaultType: UserType
  orgId: string
  user: CurrentUser
}

export const UpdateUserTypeForm: React.FC<props> = ({
  defaultType,
  orgId,
  user,
}) => {
  const [input, setInput] = useState<OrgMemberUpdateInput>({
    org_id: orgId,
    user_id: user.id,
    user_type: defaultType,
  })
  const snack = useSnackbar()

  const onSubmit = async () => {
    const res = await fetchUpdateMemberUserType(input)
    const ret: OrgInviteResponseBody = await res.json()

    switch (res.status) {
      case 200:
        mutate(`${GetPath.ORG_MEMBERS_LIST}?id=${orgId}`)
        mutate(`${GetPath.ORG_MEMBER_TYPE}?id=${user?.id}&orgId=${orgId}`)
        snack.enqueueSnackbar('Success to update user type', {
          variant: 'success',
        })
        break
      default:
        snack.enqueueSnackbar(ret.error, { variant: 'error' })
        break
    }
  }

  return (
    <Box>
      <Box component="form">
        <Box>
          <Typography variant="h6">Update User Type</Typography>
        </Box>
        <Box mt={2}>
          <Select
            label="type"
            defaultValue={defaultType}
            name="type"
            fullWidth
            onChange={(e: SelectChangeEvent) => {
              setInput({ ...input, user_type: e.target.value })
            }}
          >
            {UserTypesAll.map((t, i) => (
              <MenuItem key={i} value={t}>
                {t}
              </MenuItem>
            ))}
          </Select>
        </Box>
        <Box mt={2} textAlign="right">
          <Button
            type="button"
            variant="contained"
            disabled={defaultType === input.user_type}
            onClick={onSubmit}
          >
            Update
          </Button>
        </Box>
      </Box>
    </Box>
  )
}
