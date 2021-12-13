import { Box, Card } from '@mui/material'
import React from 'react'
import useSWR from 'swr'
import { ApiUrl } from '../../../config/env'
import { kvsResponseBody } from '../../../http/body/kv'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'

export type KvListProps = {
  projectId: string
}

export const KvList: React.FC<KvListProps> = ({ projectId }) => {
  console.log(`${ApiUrl}${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)
  const { data, error } = useSWR<kvsResponseBody, ErrorConstructor>(
    `${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  return (
    <Box>
      {data &&
        data.data &&
        data.data.map((kv, i) => (
          <Card key={i}>
            {kv.kv_key} = {kv.kv_value}
          </Card>
        ))}
    </Box>
  )
}
