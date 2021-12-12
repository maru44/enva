import { Box, Card } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { NextPage } from 'next'
import useSWR from 'swr'
import { projectsResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { GetPath } from '../../../http/fetcher'
import Link from 'next/link'

const ProjectList: NextPage<PageProps> = (props) => {
  const { data, error } = useSWR<projectsResponseBody, ErrorConstructor>(
    GetPath.PROJECT_LIST_USER,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  return (
    <Box m={2}>
      <Box>
        {data &&
          data.data &&
          data.data.map((p, i) => (
            <Card key={i}>
              {p.name} ::: {p.slug}
              <Link as={`/project/${p.slug}`} href={`/project/[slug]`} passHref>
                <a>Link</a>
              </Link>
            </Card>
          ))}
        {data && data.error && <Box>{data.error}</Box>}
        {!data && <Box>...Loading</Box>}
      </Box>
    </Box>
  )
}

const useStyle = makeStyles(() => ({
  root: {
    // padding: theme.spacing(1),
  },
}))

export default ProjectList
