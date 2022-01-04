import { ArrowBack } from '@material-ui/icons'
import { Box, IconButton, Typography } from '@mui/material'
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
  const orgId = router.query.orgId as string

  const { data, error } = useSWR<projectResponseBody, ErrorConstructor>(
    `${GetPath.PROJECT_DETAIl}?slug=${slug}&orgId=${orgId ?? ''}`,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  useRequireLogin()

  return (
    <Box mt={6}>
      {data && data.data && (
        <Box>
          <Box display="flex" flexDirection="row" alignItems="center">
            <Box mr={2}>
              <IconButton
                onClick={() => {
                  router.back()
                }}
              >
                <ArrowBack />
              </IconButton>
            </Box>
            <Typography variant="h5">{data.data.name}</Typography>
          </Box>
          <KvList projectId={data.data.id} />
        </Box>
      )}
    </Box>
  )
}

export default ProjectDetail
