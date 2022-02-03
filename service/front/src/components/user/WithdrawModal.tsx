import { Box, Button, Dialog, Typography } from '@mui/material'
import { useRouter } from 'next/router'
import { useSnackbar } from 'notistack'
import { logoutUrl } from '../../../config/aws'
import { userCRUDResponseBody } from '../../../http/body/user'
import { fetchWithdraw } from '../../../http/user'

type props = {
  onClose: () => void
  isOpen: boolean
}

export const WithdrawModal: React.FC<props> = ({ onClose, isOpen }) => {
  const snack = useSnackbar()
  const router = useRouter()

  const withdraw = async () => {
    const res = await fetchWithdraw()
    const ret: userCRUDResponseBody = await res.json()
    if (res.status === 200) {
      snack.enqueueSnackbar('Thank you. See you soon!', { variant: 'success' })
      router.push(logoutUrl)
    } else {
      snack.enqueueSnackbar(ret.error, { variant: 'error' })
    }
  }

  return (
    <Dialog onClose={onClose} open={isOpen}>
      <Box p={3}>
        <Typography variant="h5">Are you sure to withdraw?</Typography>
        <Box
          mt={2}
          display="flex"
          alignItems="center"
          justifyContent="space-between"
        >
          <Box>
            <Button
              type="button"
              variant="contained"
              onClick={onClose}
              className="subButton"
            >
              Close
            </Button>
          </Box>
          <Box>
            <Button
              type="button"
              variant="contained"
              onClick={withdraw}
              color="warning"
            >
              Withdraw
            </Button>
          </Box>
        </Box>
      </Box>
    </Dialog>
  )
}
