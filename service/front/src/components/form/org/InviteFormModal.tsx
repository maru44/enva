import { Box, Button, Dialog, Grid, TextField, Typography } from '@mui/material'
import { useSnackbar } from 'notistack'

type props = {
  orgId: string
  orgName: string
  isOpen: boolean
  onClose: () => void
}

export const InviteFormModal: React.FC<props> = ({
  orgName,
  orgId,
  isOpen,
  onClose,
}) => {
  const snack = useSnackbar()

  const posting = async () => {}

  return (
    <Dialog onClose={onClose} open={isOpen}>
      <Grid container className="dialogContainer" p={3}>
        <Grid sm={2} item />
        <Grid xs={8} item>
          <Box mt={2}>
            <Typography variant="h5">Inviting to {orgName}</Typography>
          </Box>
          <Box mt={4}>
            <TextField fullWidth label="email" type="email"></TextField>
          </Box>
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
                onClick={() => {
                  console.log(orgId)
                }}
                type="button"
                variant="contained"
              >
                Invite
              </Button>
            </Box>
          </Box>
        </Grid>
      </Grid>
    </Dialog>
  )
}
