import { Box } from '@mui/material'
import React from 'react'
import useSWR from 'swr'
import { kvsResponseBody } from '../../../http/body/kv'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { ErrorComponent } from '../error/ErrorComponent'
import { KvListTable } from './KvListTable'

export type KvListProps = {
  projectId: string
}

export const KvList: React.FC<KvListProps> = ({ projectId }) => {
  const { data, error } = useSWR<kvsResponseBody, ErrorConstructor>(
    `${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`,
    fetcherGetFromApiUrl
  )

  if (error) return <ErrorComponent />
  if (data?.error) return <ErrorComponent errBody={data} />

  return (
    <Box>
      {data && (
        <KvListTable kvs={data.data} projectId={projectId}></KvListTable>
      )}
    </Box>
  )
}
