import { Box, Typography } from '@mui/material'
import { NextPage } from 'next'
import { PageProps } from '../../../types/page'
import privacies from '../../components/static/rule/privacy.json'

type privacyJson = {
  contents: string[]
  date: string
}

const PrivacyPage: NextPage<PageProps> = (props) => {
  const ps: privacyJson = privacies ?? null
  console.log(ps)
  return (
    <Box>
      <Box mt={4}>
        <Typography variant="h5">Privacy Policy</Typography>
      </Box>
      <Box mt={4}>
        {ps &&
          ps.contents &&
          ps.contents.map((c, i) => (
            <Box mb={2}>
              <Typography key={i}>{c}</Typography>
            </Box>
          ))}
      </Box>
    </Box>
  )
}

export default PrivacyPage
