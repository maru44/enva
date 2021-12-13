import { Box, Card, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { projectResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { KvCreateForm } from '../../components/form/kv/create'
import { KvList } from '../../components/kv/list'

const ProjectDetail: NextPage<PageProps> = (props) => {
  const router = useRouter()
  const slug = router.query.slug as string

  const { data, error } = useSWR<projectResponseBody, ErrorConstructor>(
    `${GetPath.PROJECT_DETAIl}?slug=${slug}`,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  return (
    <Box>
      <Box>
        {data && data.data && (
          <Box>
            <Typography variant="h2">{data.data.name}</Typography>
            <KvList projectId={data.data.id} />
            <KvCreateForm projectId={data.data.id} />
          </Box>
        )}
      </Box>
    </Box>
  )
}

export default ProjectDetail
