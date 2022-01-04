import { Apartment, ArrowBack } from '@material-ui/icons'
import { Box, Grid, IconButton, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { OrgResponseBody, OrgsResponseBody } from '../../../http/body/org'
import { projectsResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { CommonListCard } from '../../components/CommonListCard'
import styles from '../../styles/project.module.css'

const OrgDetail: NextPage<PageProps> = (props) => {
  useRequireLogin()

  const router = useRouter()
  const slug = router.query.slug as string

  const { data, error } = useSWR<OrgResponseBody, ErrorConstructor>(
    `${GetPath.ORG_DETAIL}?slug=${slug}`,
    fetcherGetFromApiUrl
  )

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
              <IconButton>
                <Apartment />
              </IconButton>
              <Typography variant="h5">{data.data.name}</Typography>
            </Box>
            {data.data.description && (
              <Typography>{data.data.description}</Typography>
            )}
          </Box>
          <Box mt={6}>
            <OrgProjects id={data.data.id} />
          </Box>
        </Box>
      )}
    </Box>
  )
}

type orgProjectsProps = {
  id: string
}

const OrgProjects: React.FC<orgProjectsProps> = ({ id }) => {
  const { data, error } = useSWR<projectsResponseBody, ErrorConstructor>(
    `${GetPath.PROJECT_LIST_ORG}?id=${id}`,
    fetcherGetFromApiUrl
  )

  if (error) return <div></div>

  console.log(data)

  return (
    <Grid container rowSpacing={2} columnSpacing={2}>
      {data &&
        data.data &&
        data.data.map((p, i) => (
          <CommonListCard
            info={p}
            key={i}
            linkAs={`/project/${p.slug}?orgId=${id}`}
            linkHref="/project/[slug]"
            styles={styles}
          />
        ))}
    </Grid>
  )
}

export default OrgDetail
