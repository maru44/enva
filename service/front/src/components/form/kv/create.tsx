import { Box, Button, TextField } from '@mui/material'
import { makeStyles } from '@mui/styles'
import clsx from 'clsx'
import React, { useState } from 'react'
import { useSWRConfig } from 'swr'
import { kvCreateResponseBody } from '../../../../http/body/kv'
import { fetcher, GetPath } from '../../../../http/fetcher'
import { fetchCreateKv } from '../../../../http/kv'
import { KvInput } from '../../../../types/kv'

export type KvCreateProps = {
  projectId: string
}

export const KvCreateForm = ({ projectId }: KvCreateProps) => {
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
    const res = await fetcher(fetchCreateKv(input))
    const ret: kvCreateResponseBody = await res.json()
    if (res.status === 200) {
      const id = ret['data']
      console.log(id) // @TODO fix
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
          <TextField
            name="kv_key"
            variant="outlined"
            label="key"
            required
            className={clsx(classes.textField)}
          />
          <TextField
            name="kv_value"
            label="value"
            variant="outlined"
            className={clsx(classes.textField)}
          />
          <Button type="submit" variant="outlined">
            Create
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
