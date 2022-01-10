import { ArrowBack } from '@material-ui/icons'
import { Box, IconButton, Typography } from '@mui/material'
import { NextPage } from 'next'
import Link from 'next/link'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { projectResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { KvList } from '../../components/kv/KvList'

const ProjectDetail: NextPage<PageProps> = (props) => {
  useRequireLogin()
  const router = useRouter()
  const slugs = router.query.slug as string[]

  let url = GetPath.PROJECT_DETAIl as string
  if (slugs) {
    switch (slugs.length) {
      case 2:
        url = `${GetPath.PROJECT_DETAIl}?slug=${slugs[1]}&orgSlug=${slugs[0]}`
        break
      default:
        url = `${GetPath.PROJECT_DETAIl}?slug=${slugs[0]}`
    }
  }

  const { data, error } = useSWR<projectResponseBody, ErrorConstructor>(
    url,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

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
            <Box display="flex" flexDirection="row" alignItems="center">
              {data.data.org && (
                <Box display="flex" flexDirection="row" alignItems="center">
                  <Link
                    href="/org/[slug]"
                    as={`/org/${data.data.org.slug}`}
                    passHref
                  >
                    <a>
                      <Typography variant="h5">{data.data.org.name}</Typography>
                    </a>
                  </Link>
                  <Typography variant="h5"> / </Typography>
                </Box>
              )}
              <Typography variant="h5">{data.data.name}</Typography>
            </Box>
          </Box>
          <Box pl={1} pr={1} mt={4}>
            {data.data.description && (
              <Box>
                <Typography>{data.data.description}</Typography>
              </Box>
            )}
            <Box mt={4}>
              <KvList projectId={data.data.id} />
            </Box>
          </Box>
        </Box>
      )}
    </Box>
  )
}

export default ProjectDetail
