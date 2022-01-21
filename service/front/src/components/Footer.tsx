import { Box, Grid, Typography } from '@mui/material'
import { useState } from 'react'
import { TermsOfUse } from './static/rule/TermsOfUse'

export const Footer: React.FC = () => {
  const [isTermsOpen, setIsTermsOpen] = useState<boolean>(false)
  return (
    <Box>
      <Box
        display="flex"
        alignItems="center"
        justifyContent="space-between"
        pt={1}
        pb={1}
        pl={2}
        pr={1}
        mt={2}
      >
        <Box>
          <Typography>&copy; 2022 maru</Typography>
        </Box>
        <Box>
          <Typography
            onClick={() => {
              setIsTermsOpen(true)
            }}
            className="likeLink"
          >
            terms of use
          </Typography>
        </Box>
      </Box>
      <TermsOfUse
        onClose={() => {
          setIsTermsOpen(false)
        }}
        isOpen={isTermsOpen}
      />
    </Box>
  )
}
