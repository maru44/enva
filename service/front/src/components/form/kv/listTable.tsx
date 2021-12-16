import {
  Button,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material'
import React, { useState } from 'react'
import { useSWRConfig } from 'swr'
import { kvDeleteResponseBody } from '../../../../http/body/kv'
import { GetPath } from '../../../../http/fetcher'
import { fetchDeleteKv } from '../../../../http/kv'
import { Kv } from '../../../../types/kv'
import { sortKvs } from '../../../../utils/kv'
import { KvUpdateForm } from './update'

type props = {
  kvs: Kv[]
  projectId: string
}

export const KvListTable: React.FC<props> = ({ kvs, projectId }: props) => {
  const { mutate } = useSWRConfig()
  const [isOpenUpdate, setIsOpenUpdate] = useState<boolean>(false)
  const [updateKey, setUpdateKey] = useState<string>('')
  const [defaultValue, setDefaultValue] = useState<string>('')

  // delete function
  const delKeyFunc = async (keyId: string, projectId: string) => {
    try {
      const res = await fetchDeleteKv(keyId, projectId)
      const ret: kvDeleteResponseBody = await res.json()

      switch (res.status) {
        case 200:
          mutate(`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)
        default:
      }
    } catch (e) {
      // @TODO 500
      console.log(e)
    }
  }

  // close update form
  const closeUpdateForm = () => setIsOpenUpdate(false)

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
                      delKeyFunc(kv.id, projectId)
                    }}
                  >
                    Delete
                  </Button>
                  <Button
                    type="button"
                    onClick={() => {
                      setUpdateKey(kv.kv_key)
                      setDefaultValue(kv.kv_value)
                      setIsOpenUpdate(true)
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
        kvKey={updateKey}
        kvValue={defaultValue}
        projectId={projectId}
        isOpen={isOpenUpdate}
        onClose={closeUpdateForm}
      />
    </TableContainer>
  )
}
