import {
  Box,
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from '@mui/material'
import React, { useReducer } from 'react'
import {
  initialKvListState,
  kvListReducer,
} from '../../../hooks/kvs/useListTable'
import { Kv } from '../../../types/kv'
import { sortKvs } from '../../../utils/kv'
import { KvCreateTableRow } from '../form/kv/KvCreateTableRow'
import { KvUpdateModal } from '../form/kv/KvUpdateModal'
import { Delete, Edit } from '@material-ui/icons'
import { DeleteModal } from '../DeleteModal'
import { GetPath } from '../../../http/fetcher'
import styles from '../../styles/kv.module.css'

type props = {
  kvs: Kv[]
  projectId: string
}

export const KvListTable: React.FC<props> = ({ kvs, projectId }: props) => {
  const [state, dispatch] = useReducer(kvListReducer, initialKvListState)

  return (
    <Box>
      <TableContainer component={Paper} variant="outlined">
        <Table aria-label="key value sets">
          <TableHead>
            <TableRow>
              <TableCell width="30%">
                <Typography variant="subtitle1">Key</Typography>
              </TableCell>
              <TableCell width="70%">
                <Typography variant="subtitle1">Value</Typography>
              </TableCell>
              <TableCell width={128}></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {kvs &&
              sortKvs(kvs).map((kv, i) => (
                <TableRow key={i}>
                  <TableCell>
                    <Typography className={styles.breakCell} variant="inherit">
                      {kv.kv_key}
                    </Typography>
                  </TableCell>
                  <TableCell>
                    <Typography className={styles.breakCell} variant="inherit">
                      {kv.kv_value}
                    </Typography>
                  </TableCell>
                  <TableCell>
                    <Box display="flex" flexDirection="row">
                      <Box>
                        <IconButton
                          onClick={() => {
                            dispatch({
                              type: 'openUpdate',
                              targetKey: kv.kv_key,
                              updateDefaultValue: kv.kv_value,
                            })
                          }}
                        >
                          <Edit />
                        </IconButton>
                      </Box>
                      <Box ml={2}>
                        <IconButton>
                          <Delete
                            onClick={() => {
                              dispatch({
                                type: 'openDelete',
                                targetKey: kv.kv_key,
                                deleteId: kv.id,
                              })
                            }}
                          />
                        </IconButton>
                      </Box>
                    </Box>
                  </TableCell>
                </TableRow>
              ))}
            <KvCreateTableRow projectId={projectId} />
          </TableBody>
        </Table>
        <KvUpdateModal
          kvKey={state.targetKey}
          kvValue={state.updateDefaultValue}
          projectId={projectId}
          isOpen={state.isOpenUpdate}
          onClose={() => dispatch({ type: 'closeUpdate' })}
        />
        <DeleteModal
          url={`${GetPath.KV_DELETE}?kvId=${state.deleteId}&projectId=${projectId}`}
          mutateKey={`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`}
          isOpen={state.isOpenDelete}
          onClose={() => dispatch({ type: 'closeDelete' })}
          Message={
            <Typography variant="h5">
              Are you sure to delete <br />
              <b>{state.targetKey}</b>?
            </Typography>
          }
        />
      </TableContainer>
    </Box>
  )
}
