import {
  Box,
  Button,
  TableCell,
  TableRow,
  TextField,
  Tooltip,
} from '@mui/material'
import { useSnackbar } from 'notistack'
import React, { useState } from 'react'
import { useSWRConfig } from 'swr'
import { kvCreateResponseBody } from '../../../../http/body/kv'
import { GetPath } from '../../../../http/fetcher'
import { fetchCreateKv } from '../../../../http/kv'
import { KvInput } from '../../../../types/kv'
import { isSlug } from '../../../../utils/slug'
import styles from '../../../styles/kv.module.css'

export type KvUpsertProps = {
  projectId: string
}

// if env_key is exists >> update method is executed
export const KvCreateTableRow = ({ projectId }: KvUpsertProps) => {
  const { mutate } = useSWRConfig()
  const snack = useSnackbar()

  const [key, setKey] = useState<string>('')
  const [value, setValue] = useState<string>('')

  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const input: KvInput = {
      project_id: projectId,
      input: {
        kv_key: key,
        kv_value: value,
      },
    }
    const res = await fetchCreateKv(input)
    const ret: kvCreateResponseBody = await res.json()
    if (res.status === 200) {
      const id = ret['data']
      mutate(`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)

      setKey('')
      setValue('')
    } else {
      const message = ret['error']
      snack.enqueueSnackbar(message, { variant: 'error' })
    }
  }

  return (
    <TableRow>
      <TableCell colSpan={3} width="100%" sx={{ padding: 0 }}>
        <form onSubmit={submit}>
          <Box
            display="flex"
            flexDirection="row"
            justifyContent="space-between"
          >
            <Box width="30%" p={2}>
              <Tooltip
                title="must be slug"
                disableHoverListener
                arrow
                placement="bottom"
                open={key !== '' && !isSlug(key)}
              >
                <TextField
                  name="kv_key"
                  variant="outlined"
                  label="key"
                  required
                  fullWidth
                  value={key}
                  onChange={(e) => {
                    setKey(e.currentTarget.value)
                  }}
                />
              </Tooltip>
            </Box>
            <Box width="70%" p={2}>
              <TextField
                name="kv_value"
                label="value"
                variant="outlined"
                fullWidth
                value={value}
                onChange={(e) => {
                  setValue(e.currentTarget.value)
                }}
              />
            </Box>
            <Box flex={1} p={2}>
              <Box height={24} pt={1}>
                <Button
                  type="submit"
                  variant="contained"
                  color="primary"
                  className={styles.createButton}
                  disabled={!key || !isSlug(key)}
                >
                  CREATE
                </Button>
              </Box>
            </Box>
          </Box>
        </form>
      </TableCell>
    </TableRow>
  )
}
