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
import { makeStyles } from '@mui/styles'
import clsx from 'clsx'
import React, { useReducer } from 'react'
import {
  initialKvListState,
  kvListReducer,
} from '../../../hooks/kvs/useListTable'
import { Kv } from '../../../types/kv'
import { sortKvs } from '../../../utils/kv'
import theme from '../../theme/theme'
import { KvInsertTableRow } from '../form/kv/createTablerow'
import { KvDeleteModal } from '../form/kv/deleteModal'
import { KvUpdateForm } from '../form/kv/update'
import { Delete, Edit } from '@material-ui/icons'

type props = {
  kvs: Kv[]
  projectId: string
}

export const KvListTable: React.FC<props> = ({ kvs, projectId }: props) => {
  const [state, dispatch] = useReducer(kvListReducer, initialKvListState)

  const classes = useStyles(theme)

  return (
    <TableContainer
      component={Paper}
      className={clsx(classes.tableContainer)}
      variant="outlined"
    >
      <Table aria-label="key value sets">
        <TableHead>
          <TableRow>
            <TableCell width="30%">
              <Typography variant="subtitle1">Key</Typography>
            </TableCell>
            <TableCell width="70%">
              <Typography variant="subtitle1">Value</Typography>
            </TableCell>
            <TableCell width={100}></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {kvs &&
            sortKvs(kvs).map((kv, i) => (
              <TableRow key={i}>
                <TableCell>
                  <Typography
                    className={clsx(classes.breakCell)}
                    variant="inherit"
                  >
                    {kv.kv_key}
                  </Typography>
                </TableCell>
                <TableCell>
                  <Typography
                    className={clsx(classes.breakCell)}
                    variant="inherit"
                  >
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
          <KvInsertTableRow projectId={projectId} />
        </TableBody>
      </Table>
      <KvUpdateForm
        kvKey={state.targetKey}
        kvValue={state.updateDefaultValue}
        projectId={projectId}
        isOpen={state.isOpenUpdate}
        onClose={() => dispatch({ type: 'closeUpdate' })}
      />
      <KvDeleteModal
        kvId={state.deleteId}
        projectId={projectId}
        kvKey={state.targetKey}
        isOpen={state.isOpenDelete}
        onClose={() => dispatch({ type: 'closeDelete' })}
      ></KvDeleteModal>
    </TableContainer>
  )
}

const useStyles = makeStyles((theme) => ({
  tableContainer: {
    // backgroundColor: theme.palette.grey[200],
    marginTop: theme.spacing(2),
  },
  breakCell: {
    wordBreak: 'break-all',
  },
}))
