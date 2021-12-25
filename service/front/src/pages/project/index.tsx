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
import { ProjectListCard } from '../../components/project/ProjectListCard'

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

  return (
    <Box mt={2} width="100%">
      <Grid container rowSpacing={2} columnSpacing={2}>
        {data &&
          data.data &&
          data.data.map((p, i) => (
            <ProjectListCard
              project={p}
              key={i}
              startDeleteFunc={() =>
                dispatch({
                  type: 'openDelete',
                  targetKey: p.name,
                  deleteId: p.id,
                })
              }
            />
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

export default ProjectList
