import { AccountCircle } from '@material-ui/icons'
import { Box, Icon, Tooltip, Typography } from '@mui/material'
import useSWR from 'swr'
import { OrgMemberListResponseBody } from '../../../../http/body/org'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'
import { UserType, UserTypesAll } from '../../../../types/user'

type props = {
  id: string
}

export const MembersList: React.FC<props> = ({ id }) => {
  const { data, error } = useSWR<OrgMemberListResponseBody>(
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
              <Box mb={1}>
                <Typography>{type}</Typography>
              </Box>
              {data.data[type].map((u, i) => (
                <Box key={`${type}_${i}`}>
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
    </Box>
  )
}
