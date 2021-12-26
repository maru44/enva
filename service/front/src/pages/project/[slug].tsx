import { Box, Card, Container, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { projectResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { KvList } from '../../components/kv/KvList'

const ProjectDetail: NextPage<PageProps> = (props) => {
  const router = useRouter()
  const slug = router.query.slug as string

  const { data, error } = useSWR<projectResponseBody, ErrorConstructor>(
    `${GetPath.PROJECT_DETAIl}?slug=${slug}`,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  useRequireLogin()

  return (
    <Container>
      {data && data.data && (
        <Box mt={4}>
          <Typography variant="h5">{data.data.name}</Typography>
          <KvList projectId={data.data.id} />
        </Box>
      )}
    </Container>
  )
}

export default ProjectDetail
