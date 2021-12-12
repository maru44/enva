import { Box, Card } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { projectResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'

const ProjectDetail: NextPage<PageProps> = (props) => {
  const router = useRouter()
  const slug = router.query.slug

  const { data, error } = useSWR<projectResponseBody, ErrorConstructor>(
    `${GetPath.PROJECT_DETAIl}?slug=${slug}`,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  return (
    <Box>
      <Box>
        {data &&
          data.data &&
          <Box>{data.data.name}</Box>}
      </Box>
    </Box>
  )
}

export default ProjectDetail
