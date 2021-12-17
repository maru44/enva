import { Box, Button, Dialog, DialogTitle, TextField } from '@mui/material'
import React, { useState } from 'react'
import { mutate } from 'swr'
import { GetPath } from '../../../../http/fetcher'
import { fetchUpdateKv } from '../../../../http/kv'
import { KvInput } from '../../../../types/kv'

type props = {
  kvKey: string
  kvValue: string
  projectId: string
  isOpen: boolean
  //   onClose: React.Dispatch<React.SetStateAction<boolean>>
  onClose: () => void
}

export const KvUpdateForm: React.FC<props> = ({
  kvKey,
  kvValue,
  projectId,
  isOpen,
  onClose,
}) => {
  const [val, setVal] = useState<string>(kvValue)

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setVal(e.currentTarget.value)
  }

  const update = async () => {
    try {
      const input: KvInput = {
        project_id: projectId,
        input: {
          kv_key: kvKey,
          kv_value: val,
        },
      }
      const res = await fetchUpdateKv(input)
      const ret = await res.json()
      switch (res.status) {
        case 200:
          // @TODO success alert
          mutate(`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)
        default:
      }

      onClose()
    } catch (e) {
      // @TODO alert 500
      console.log(e)
    }
  }

  return (
    <Dialog onClose={onClose} open={isOpen}>
      <Box m={3}>
        <DialogTitle>Edit: {kvKey}</DialogTitle>
        <TextField
          type="text"
          defaultValue={kvValue}
          placeholder="value"
          onChange={onChange}
        />
        <Button onClick={update} type="button">
          Update
        </Button>
      </Box>
    </Dialog>
  )
}
