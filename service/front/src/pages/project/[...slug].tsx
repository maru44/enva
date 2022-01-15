import { ArrowBack, FileCopy } from '@material-ui/icons'
import { Box, IconButton, Tooltip, Typography } from '@mui/material'
import { NextPage } from 'next'
import Link from 'next/link'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { projectResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { ErrorComponent } from '../../components/error/ErrorComponent'
import { KvList } from '../../components/kv/KvList'
import styles from '../../styles/project.module.css'

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

  if (error) return <ErrorComponent />
  if (data?.error) return <ErrorComponent errBody={data} />

  const sampleJson = `{
  "env_file_name": ".envrc",
  "projectSlug": "${data?.data.slug}",${
    data?.data.org
      ? `
  "org_slug": "${data?.data.org.slug}",`
      : ''
  }
  "pre_sentence": "# this is optional value\\n# you can write any thing",
  "suf_sentence": "# this is optional value\\n# Have a nice day!"
}`

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
          <Box mt={6} pl={1} pr={1}>
            <Typography variant="h6">
              Sample enva.json for this project
            </Typography>
            <Box mt={2}>
              <Box className={styles.sampleJson} pb={2} pr={3} pl={3}>
                <Box textAlign="right" pt={1}>
                  <Tooltip title="copy" arrow>
                    <IconButton
                      onClick={() => {
                        navigator.clipboard.writeText(sampleJson)
                      }}
                    >
                      <FileCopy />
                    </IconButton>
                  </Tooltip>
                </Box>
                <Typography>
                  <pre>
                    <code>{sampleJson}</code>
                  </pre>
                </Typography>
              </Box>
            </Box>
          </Box>
        </Box>
      )}
    </Box>
  )
}

export default ProjectDetail
