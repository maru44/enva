import { Box, Button, TextField } from '@mui/material'
import { makeStyles } from '@mui/styles'
import clsx from 'clsx'
import React from 'react'
import { useSWRConfig } from 'swr'
import { kvCreateResponseBody } from '../../../../http/body/kv'
import { fetcher, GetPath } from '../../../../http/fetcher'
import { fetchCreateKv, fetchUpdateKv } from '../../../../http/kv'
import { KvInput } from '../../../../types/kv'

export type KvUpsertProps = {
  projectId: string
  env_key?: string
}

// if env_key is exists >> update method is executed
export const KvUpsertForm = ({ projectId, env_key }: KvUpsertProps) => {
  const { mutate } = useSWRConfig()
  
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
    const res = env_key
      ? await fetcher(fetchUpdateKv(input))
      : await fetcher(fetchCreateKv(input))
    const ret: kvCreateResponseBody = await res.json()
    if (res.status === 200) {
      const id = ret['data']
      mutate(`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)
    } else {
      const message = ret['error']
      console.log(message) // @TODO fix
    }
  }

  const classes = useStyle()

  return (
    <Box>
      <form onSubmit={submit}>
        <Box display="flex" flexDirection="column">
          {!env_key && (
            <TextField
              name="kv_key"
              variant="outlined"
              label="key"
              required
              className={clsx(classes.textField)}
            />
          )}
          {env_key && (
            <TextField
              name="kv_key"
              variant="outlined"
              label="key"
              type="hidden"
              value={env_key}
              required
              className={clsx(classes.textField)}
            />
          )}
          <TextField
            name="kv_value"
            label="value"
            variant="outlined"
            className={clsx(classes.textField)}
          />
          <Button type="submit" variant="outlined">
            {env_key ? 'UPDATE' : 'Create'}
          </Button>
        </Box>
      </form>
    </Box>
  )
}

const useStyle = makeStyles(() => ({
  textField: {
    marginBottom: 8,
  },
}))
