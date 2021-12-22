import { Box, Card, Container, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { projectResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { KvUpsertForm } from '../../components/form/kv/create'
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
    <Container>
      {data && data.data && (
        <Box>
          <Typography variant="h4">{data.data.name}</Typography>
          <KvList projectId={data.data.id} />
          <KvUpsertForm projectId={data.data.id} />
        </Box>
      )}
    </Container>
  )
}

export default ProjectDetail
