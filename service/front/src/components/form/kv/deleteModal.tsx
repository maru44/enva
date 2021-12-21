import { Box, Button, Dialog, DialogTitle, Grid } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { mutate } from 'swr'
import { kvDeleteResponseBody } from '../../../../http/body/kv'
import { GetPath } from '../../../../http/fetcher'
import { fetchDeleteKv } from '../../../../http/kv'
import theme from '../../../theme/theme'

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
    onClose()
  }

  const classes = useStyles(theme)

  return (
    <Dialog onClose={onClose} open={isOpen}>
      <Grid container className={classes.dialogContainer} p={3}>
        <Grid sm={2} />
        <Grid lg={8} sm={8}>
          <DialogTitle>
            Are you sure to delete <br />
            <b>{kvKey}</b>?
          </DialogTitle>
          <Grid mt={2} justifyContent="space-between" container>
            <Button
              onClick={onClose}
              type="button"
              variant="contained"
              className={classes.subButton}
            >
              Close
            </Button>
            <Button onClick={onDelete} type="button" variant="contained">
              Delete
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
