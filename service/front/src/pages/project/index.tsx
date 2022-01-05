import { Box, Grid, Typography } from '@mui/material'
import { NextPage } from 'next'
import useSWR from 'swr'
import { projectsResponseBody } from '../../../http/body/project'
import { fetcherGetFromApiUrl } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'
import { GetPath } from '../../../http/fetcher'
import { DeleteModal } from '../../components/DeleteModal'
import {
  initialProjectListState,
  projectListReducer,
} from '../../../hooks/kvs/useListProject'
import { useReducer } from 'react'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { CommonListCard } from '../../components/CommonListCard'
import styles from '../../styles/project.module.css'

const ProjectList: NextPage<PageProps> = (props) => {
  useRequireLogin()

  const { data, error } = useSWR<projectsResponseBody, ErrorConstructor>(
    GetPath.PROJECT_LIST,
    fetcherGetFromApiUrl
  )
  const [state, dispatch] = useReducer(
    projectListReducer,
    initialProjectListState
  )

  // @TODO error handling
  if (error) console.log(error)

  return (
    <Box mt={6} width="100%">
      <Typography variant="h5">Projects</Typography>
      <Grid container mt={1} rowSpacing={2} columnSpacing={2}>
        {data &&
          data.data &&
          data.data.map((p, i) => (
            <CommonListCard
              info={p}
              key={i}
              startDeleteFunc={() => {
                dispatch({
                  type: 'openDelete',
                  targetKey: p.name,
                  deleteId: p.id,
                })
              }}
              linkAs={
                p.org ? `/project/${p.org.id}/${p.slug}` : `/project/${p.slug}`
              }
              linkHref="/project/[...slug]"
              styles={styles}
              name={p.org && `${p.org.slug}/${p.name}`}
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
