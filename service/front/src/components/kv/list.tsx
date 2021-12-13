import { Box, Card } from '@mui/material'
import React from 'react'
import useSWR from 'swr'
import { ApiUrl } from '../../../config/env'
import { kvsResponseBody } from '../../../http/body/kv'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { sortKvs } from '../../../utils/kv'
import { KvUpsertForm } from '../form/kv/create'

export type KvListProps = {
  projectId: string
}

export const KvList: React.FC<KvListProps> = ({ projectId }) => {
  const { data, error } = useSWR<kvsResponseBody, ErrorConstructor>(
    `${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  return (
    <Box>
      {data &&
        data.data &&
        sortKvs(data.data).map((kv, i) => (
          <Card key={i}>
            {kv.kv_key} = {kv.kv_value}
            <KvUpsertForm projectId={projectId} env_key={kv.kv_key} />
          </Card>
        ))}
    </Box>
  )
}
