import {
  Box,
  Button,
  Dialog,
  FormControl,
  Grid,
  TextField,
  Typography,
} from '@mui/material'
import { useSnackbar } from 'notistack'
import React, { useState } from 'react'
import { mutate } from 'swr'
import { kvCreateResponseBody } from '../../../../http/body/kv'
import { GetPath } from '../../../../http/fetcher'
import { fetchUpdateKv } from '../../../../http/kv'
import { KvInput } from '../../../../types/kv'

type props = {
  kvKey: string
  kvValue: string
  projectId: string
  isOpen: boolean
  onClose: () => void
}

export const KvUpdateModal: React.FC<props> = ({
  kvKey,
  kvValue,
  projectId,
  isOpen,
  onClose,
}) => {
  const [val, setVal] = useState<string>(kvValue)
  const snack = useSnackbar()

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
      const ret: kvCreateResponseBody = await res.json()
      switch (res.status) {
        case 200:
          snack.enqueueSnackbar('succeeded to update value', {
            variant: 'success',
          })
          mutate(`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)
          break
        default:
          snack.enqueueSnackbar(ret.error, { variant: 'error' })
          break
      }
      onClose()
    } catch (e) {
      snack.enqueueSnackbar('Internal Server Error', { variant: 'error' })
    }
  }

  return (
    <Dialog
      onClose={() => {
        onClose()
        setVal('')
      }}
      open={isOpen}
    >
      <Grid container className="dialogContainer" p={3}>
        <Grid item sm={2} />
        <Grid item lg={8} sm={8}>
          <Typography variant="h5">Edit: {kvKey}</Typography>
          <Box mt={2}>
            <FormControl fullWidth>
              <TextField
                type="text"
                defaultValue={kvValue}
                placeholder="value"
                onChange={onChange}
              />
            </FormControl>
          </Box>
          <Grid container justifyContent="space-between" mt={2}>
            <Button
              onClick={() => {
                onClose()
                setVal('')
              }}
              variant="contained"
              type="button"
              className="subButton"
            >
              Close
            </Button>
            <Button
              onClick={update}
              variant="contained"
              type="button"
              disabled={!val || kvValue === val}
            >
              Update
            </Button>
          </Grid>
        </Grid>
      </Grid>
    </Dialog>
  )
}
