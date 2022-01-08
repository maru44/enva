import { Box } from '@mui/material'
import useSWR from 'swr'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'

type props = {
  id: string
}

export const MembersList: React.FC<props> = ({ id }) => {
  const { data, error } = useSWR(
    `${GetPath.ORG_MEMBERS_LIST}?id=${id}`,
    fetcherGetFromApiUrl
  )

  console.log(data)

  return <Box></Box>
}
