import { Box, Typography } from '@mui/material'
import { termly } from './termly'

export const TermsOfUse: React.FC = () => {
  return (
    <Box>
      <Box mt={2}>
        <Typography variant="h4">Terms of use</Typography>
      </Box>
      <Box mt={4}>
        <div dangerouslySetInnerHTML={{ __html: termly }} />
      </Box>
    </Box>
  )
}
