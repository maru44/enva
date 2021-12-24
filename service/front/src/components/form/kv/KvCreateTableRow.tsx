import { Box, Button, TableCell, TableRow, TextField } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { useSnackbar } from 'notistack'
import React, { useRef } from 'react'
import { useSWRConfig } from 'swr'
import { kvCreateResponseBody } from '../../../../http/body/kv'
import { fetcher, GetPath } from '../../../../http/fetcher'
import { fetchCreateKv } from '../../../../http/kv'
import { KvInput } from '../../../../types/kv'

export type KvUpsertProps = {
  projectId: string
}

// if env_key is exists >> update method is executed
export const KvCreateTableRow = ({ projectId }: KvUpsertProps) => {
  const { mutate } = useSWRConfig()
  const snack = useSnackbar()

  const refKey = useRef<HTMLInputElement>(null)
  const refValue = useRef<HTMLInputElement>(null)

  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const t = e.currentTarget
    const k = t.kv_key.value
    const v = t.kv_value.value

    const input: KvInput = {
      project_id: projectId,
      input: {
        kv_key: k,
        kv_value: v,
      },
    }
    const res = await fetcher(fetchCreateKv(input))
    const ret: kvCreateResponseBody = await res.json()
    if (res.status === 200) {
      const id = ret['data']
      mutate(`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)

      if (refKey.current) {
        refKey.current.value = ''
      }
      if (refValue.current) {
        refValue.current.value = ''
      }
    } else {
      const message = ret['error']
      snack.enqueueSnackbar(message, { variant: 'error' })
    }
  }

  const classes = useStyles()

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
              <TextField
                name="kv_key"
                variant="outlined"
                label="key"
                required
                fullWidth
                inputRef={refKey}
              />
            </Box>
            <Box width="70%" p={2}>
              <TextField
                name="kv_value"
                label="value"
                variant="outlined"
                fullWidth
                inputRef={refValue}
              />
            </Box>
            <Box flex={1} p={2}>
              <Box height={24} pt={1}>
                <Button
                  type="submit"
                  variant="contained"
                  color="success"
                  className={classes.createButton}
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

const useStyles = makeStyles(() => ({
  createButton: {
    width: 96,
  },
}))
