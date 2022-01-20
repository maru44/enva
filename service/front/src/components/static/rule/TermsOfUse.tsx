import { Box, Dialog, Typography } from '@mui/material'
import { termly } from './termly'

export const TermsOfUse: React.FC<{ onClose: () => void; isOpen: boolean }> = ({
  onClose,
  isOpen,
}) => {
  return (
    <Dialog onClose={onClose} open={isOpen} fullWidth maxWidth="md">
      <Box p={3}>
        <div dangerouslySetInnerHTML={{ __html: termly }} />
      </Box>
    </Dialog>
  )
}
