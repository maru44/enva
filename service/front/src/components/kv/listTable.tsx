import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material'
import React, { useReducer } from 'react'
import {
  initialKvListState,
  kvListReducer,
} from '../../../hooks/kvs/useListTable'
import { Kv } from '../../../types/kv'
import { sortKvs } from '../../../utils/kv'
import { KvDeleteModal } from '../form/kv/deleteModal'
import { KvUpdateForm } from '../form/kv/update'

type props = {
  kvs: Kv[]
  projectId: string
}

export const KvListTable: React.FC<props> = ({ kvs, projectId }: props) => {
  const [state, dispatch] = useReducer(kvListReducer, initialKvListState)

  return (
    <TableContainer>
      <Table aria-label="key value sets">
        <TableHead>
          <TableRow>
            <TableCell>Key</TableCell>
            <TableCell>Value</TableCell>
            <TableCell>Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {kvs &&
            sortKvs(kvs).map((kv, i) => (
              <TableRow key={i}>
                <TableCell>{kv.kv_key}</TableCell>
                <TableCell>{kv.kv_value}</TableCell>
                <TableCell>
                  <Button
                    type="button"
                    onClick={() => {
                      dispatch({
                        type: 'openDelete',
                        targetKey: kv.kv_key,
                        deleteId: kv.id,
                      })
                    }}
                  >
                    Delete
                  </Button>
                  <Button
                    type="button"
                    onClick={() => {
                      dispatch({
                        type: 'openUpdate',
                        targetKey: kv.kv_key,
                        updateDefaultValue: kv.kv_value,
                      })
                    }}
                  >
                    Edit
                  </Button>
                </TableCell>
              </TableRow>
            ))}
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
