import { AccountCircle } from '@material-ui/icons'
import { Box, Icon, Tooltip, Typography } from '@mui/material'
import { useState } from 'react'
import useSWR from 'swr'
import { useRequireLogin } from '../../../../hooks/useRequireLogin'
import { OrgMemberListResponseBody } from '../../../../http/body/org'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'
import { UserTypeDescription } from '../../../../types/org'
import { CurrentUser, UserType, UserTypesAll } from '../../../../types/user'
import { MemberDetailModal } from './MemberDetailModal'

type props = {
  id: string
  currentUserType: UserType
}

export const MembersList: React.FC<props> = ({ id, currentUserType }) => {
  useRequireLogin()
  const [selectedMember, setSelectedMember] = useState<CurrentUser | undefined>(
    undefined
  )
  const [membersType, setMembersType] = useState<UserType | undefined>(
    undefined
  )
  const { data, error } = useSWR<OrgMemberListResponseBody, ErrorConstructor>(
    `${GetPath.ORG_MEMBERS_LIST}?id=${id}`,
    fetcherGetFromApiUrl
  )

  if (!data || !data.data) {
    return <div></div>
  }

  return (
    <Box>
      {UserTypesAll.map(
        (type) =>
          data.data[type] && (
            <Box key={type} mb={2}>
              <Box mb={1} display="flex" alignItems="center">
                <Box>
                  <Typography variant="subtitle1">{type}</Typography>
                </Box>
                <Box ml={2}>
                  <Typography variant="caption">
                    <small>{UserTypeDescription[type]}</small>
                  </Typography>
                </Box>
              </Box>
              {data.data[type].map((u, i) => (
                <Box
                  key={`${type}_${i}`}
                  onClick={() => {
                    setSelectedMember(u)
                    setMembersType(type)
                  }}
                >
                  <Tooltip title={u.username}>
                    <Icon>
                      <AccountCircle />
                    </Icon>
                  </Tooltip>
                </Box>
              ))}
            </Box>
          )
      )}
      <MemberDetailModal
        orgId={id}
        user={selectedMember}
        defaultType={membersType}
        currentUserType={currentUserType}
        onClose={() => {
          setSelectedMember(undefined)
        }}
      />
    </Box>
  )
}
