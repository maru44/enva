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
      <Box
        mt={4}
        display="flex"
        justifyContent="space-between"
        alignItems="center"
      >
        <Typography variant="h5">Privacy Policy</Typography>
        <Typography variant="h6">{ps?.date}</Typography>
      </Box>
      <Box mt={6}>
        {ps &&
          ps.contents &&
          ps.contents.map((c, i) => (
            <Box key={i} mb={2}>
              <Typography>{c}</Typography>
            </Box>
          ))}
      </Box>
      <Box mt={12}></Box>
    </Box>
  )
}

export default PrivacyPage
