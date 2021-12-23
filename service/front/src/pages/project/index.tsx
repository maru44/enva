import { Box, Card, Grid, Paper, Typography } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { NextPage } from 'next'
import useSWR from 'swr'
import { projectsResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { GetPath } from '../../../http/fetcher'
import Link from 'next/link'
import clsx from 'clsx'
import theme from '../../theme/theme'

const ProjectList: NextPage<PageProps> = (props) => {
  const { data, error } = useSWR<projectsResponseBody, ErrorConstructor>(
    GetPath.PROJECT_LIST_USER,
    fetcherGetFromApiUrl
  )

  if (error) console.log(error)

  const classes = useStyles(theme)

  return (
    <Box mt={2} width="100%">
      <Grid container rowSpacing={2} columnSpacing={2}>
        {data &&
          data.data &&
          data.data.map((p, i) => (
            <Grid item md={4} xs={6} key={i}>
              <Card
                className={clsx(classes.card, 'hrefBox')}
                component={Paper}
                variant="outlined"
              >
                <Box pl={2} pr={2} pt={1} pb={1}>
                  <Typography variant="h6">{p.name}</Typography>
                </Box>
                <Link
                  as={`/project/${p.slug}`}
                  href={`/project/[slug]`}
                  passHref
                >
                  <a className="hrefBoxIn"></a>
                </Link>
              </Card>
            </Grid>
          ))}
        {data && data.error && <Box>{data.error}</Box>}
        {!data && <Box>...Loading</Box>}
      </Grid>
    </Box>
  )
}

const useStyles = makeStyles((theme) => ({
  root: {
    // padding: theme.spacing(1),
  },
  card: {
    height: theme.spacing(15),
  },
}))

export default ProjectList
