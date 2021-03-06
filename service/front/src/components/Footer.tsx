import { Box, Grid, Typography } from '@mui/material'
import Link from 'next/link'
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
        <Box display="flex" alignItems="center">
          <Link passHref href="https://forms.gle/oVf98yrmcnhwc1SLA">
            <a target="_new">Inquiry</a>
          </Link>
          <Typography
            onClick={() => {
              setIsTermsOpen(true)
            }}
            className="likeLink"
            ml={2}
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
