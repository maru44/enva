import { Box, Button, Dialog, DialogTitle, Grid } from '@mui/material'
import { mutate } from 'swr'
import { kvDeleteResponseBody } from '../../../../http/body/kv'
import { GetPath } from '../../../../http/fetcher'
import { fetchDeleteKv } from '../../../../http/kv'

type props = {
  kvId: string
  kvKey: string
  projectId: string
  isOpen: boolean
  onClose: () => void
}

export const KvDeleteModal: React.FC<props> = ({
  kvId,
  kvKey,
  projectId,
  isOpen,
  onClose,
}) => {
  const onDelete = async () => {
    try {
      const res = await fetchDeleteKv(kvId, projectId)
      const ret: kvDeleteResponseBody = await res.json()

      switch (res.status) {
        case 200:
          mutate(`${GetPath.KVS_BY_PROJECT}?projectId=${projectId}`)
        default:
      }
    } catch (e) {
      // @
    }
  }

  return (
    <Dialog onClose={onClose} open={isOpen}>
      <Box m={3}>
        <DialogTitle>
          Are you sure to delete <b>{kvKey}</b>?
        </DialogTitle>
        <Grid>
          <Button onClick={onDelete} type="button">
            Delete
          </Button>
        </Grid>
      </Box>
    </Dialog>
  )
}
