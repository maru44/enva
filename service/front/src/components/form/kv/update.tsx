import {
  Box,
  Button,
  Dialog,
  DialogTitle,
  FormControl,
  Grid,
  TextField,
  Typography,
} from '@mui/material'
import React, { useState } from 'react'
import { mutate } from 'swr'
import { GetPath } from '../../../../http/fetcher'
import { fetchUpdateKv } from '../../../../http/kv'
import { KvInput } from '../../../../types/kv'
import makeStyles from '@mui/styles/makeStyles'
import theme from '../../../theme/theme'

type props = {
  kvKey: string
  kvValue: string
  projectId: string
  isOpen: boolean
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

  const classes = useStyles(theme)

  return (
    <Dialog
      onClose={() => {
        onClose()
        setVal('')
      }}
      open={isOpen}
    >
      <Grid container className={classes.dialogContainer} p={3}>
        <Grid sm={2} />
        <Grid lg={8} sm={8}>
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
              className={classes.subButton}
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

const useStyles = makeStyles((theme) => ({
  dialogContainer: {
    maxWidth: '600px',
    width: '80vw',
  },
  subButton: {
    backgroundColor: theme.palette.grey[700],
  },
}))
