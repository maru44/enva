import { Box, Typography } from '@mui/material'
import {
  errorResponseBody,
  internalServerErrorInFront,
} from '../../../http/body/error'

type props = {
  errBody?: errorResponseBody
}

export const ErrorComponent: React.FC<props> = ({ errBody }) => {
  if (!errBody) errBody = internalServerErrorInFront()

  return (
    <Box mt="36vh" display="flex" alignItems="center" justifyContent="center">
      <Box textAlign="center">
        <Typography variant="h4">{errBody.status}</Typography>
        <Typography>{errBody.error}</Typography>
      </Box>
    </Box>
  )
}
