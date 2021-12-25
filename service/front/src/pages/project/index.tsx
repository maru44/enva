import {
  Box,
  Card,
  Grid,
  IconButton,
  Paper,
  Tooltip,
  Typography,
} from '@mui/material'
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
import { Delete } from '@material-ui/icons'
import { DeleteModal } from '../../components/DeleteModal'
import {
  initialProjectListState,
  projectListReducer,
} from '../../../hooks/kvs/useListProject'
import { useReducer } from 'react'

const ProjectList: NextPage<PageProps> = (props) => {
  const { data, error } = useSWR<projectsResponseBody, ErrorConstructor>(
    GetPath.PROJECT_LIST_USER,
    fetcherGetFromApiUrl
  )
  const [state, dispatch] = useReducer(
    projectListReducer,
    initialProjectListState
  )

  // @TODO error handling
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
                <Grid container pl={2} pr={2} pt={1} pb={1}>
                  <Grid
                    item
                    xs={12}
                    display="flex"
                    flexDirection="row"
                    alignItems="center"
                    justifyContent="space-between"
                  >
                    <Grid item flex={1} overflow="hidden">
                      <Typography variant="h6">{p.name}</Typography>
                    </Grid>
                    <Grid item width={40}>
                      <Tooltip title="delete project" arrow>
                        <IconButton
                          className={classes.deleteIcon}
                          onClick={() => {
                            dispatch({
                              type: 'openDelete',
                              deleteId: p.id,
                              targetKey: p.name,
                            })
                          }}
                        >
                          <Delete />
                        </IconButton>
                      </Tooltip>
                    </Grid>
                  </Grid>
                </Grid>
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
      <DeleteModal
        url={`${GetPath.PROJECT_DELETE}?projectId=${state.deleteId}`}
        isOpen={state.isOpenDelete}
        mutateKey={GetPath.PROJECT_LIST_USER}
        Message={
          <Typography variant="h5">Delete {state.targetKey}?</Typography>
        }
        onClose={() => dispatch({ type: 'closeDelete' })}
      />
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
  deleteIcon: {
    zIndex: 100,
  },
}))

export default ProjectList
