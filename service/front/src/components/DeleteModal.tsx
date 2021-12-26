import { Box, Button, Dialog, Grid } from '@mui/material'
import { useSnackbar } from 'notistack'
import { ReactNode } from 'react'
import { mutate } from 'swr'
import { IsDevelopment } from '../../config/env'
import { deleteResponseBody } from '../../http/body/common'
import { fetchBaseApi } from '../../http/fetcher'
import theme from '../theme/theme'

export type DeleteModalProps = {
  Message: ReactNode
  isOpen: boolean
  url: string
  mutateKey?: string
  //   onSubmit: () => Promise<void>
  onClose: () => void
}

export const DeleteModal: React.FC<DeleteModalProps> = ({
  Message,
  isOpen,
  url,
  mutateKey,
  onClose,
}) => {
  const snack = useSnackbar()

  const onDelete = async () => {
    try {
      const res = await fetchBaseApi(url, 'DELETE')
      const ret: deleteResponseBody = await res.json()

      switch (res.status) {
        case 200:
          snack.enqueueSnackbar('Success to delete project', {
            variant: 'success',
          })
          mutateKey && mutate(mutateKey)
          break
        default:
          snack.enqueueSnackbar(ret.error, { variant: 'error' })
          break
      }
      onClose()
    } catch (e) {
      IsDevelopment && console.log(e)
      snack.enqueueSnackbar('Internal Server Error', { variant: 'error' })
      onClose()
    }
  }

  return (
    <Dialog onClose={onClose} open={isOpen}>
      <Grid container className="dialogContainer" p={3}>
        <Grid sm={2} />
        <Grid xs={8}>
          {Message}
          <Box mt={4} display="flex" justifyContent="space-between">
            <Box>
              <Button
                onClick={onClose}
                type="button"
                variant="contained"
                className="subButton"
              >
                Close
              </Button>
            </Box>
            <Box>
              <Button
                onClick={onDelete}
                type="button"
                variant="contained"
                color="warning"
              >
                Delete
              </Button>
            </Box>
          </Box>
        </Grid>
      </Grid>
    </Dialog>
  )
}
