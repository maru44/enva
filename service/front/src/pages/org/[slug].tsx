import { ArrowBack } from '@material-ui/icons'
import { Box, IconButton, Typography } from '@mui/material'
import { NextPage } from 'next'
import { useRouter } from 'next/router'
import useSWR from 'swr'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { OrgResponseBody, OrgsResponseBody } from '../../../http/body/org'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { PageProps } from '../../../types/page'

const OrgDetail: NextPage<PageProps> = (props) => {
  useRequireLogin()

  const router = useRouter()
  const slug = router.query.slug as string

  const { data, error } = useSWR<OrgResponseBody>(
    `${GetPath.ORG_DETAIL}?slug=${slug}`,
    fetcherGetFromApiUrl
  )

  return (
    <Box mt={6}>
      {data && data.data && (
        <Box>
          <Box display="flex" flexDirection="row">
            <Box mr={2}>
              <IconButton
                onClick={() => {
                  router.back()
                }}
              >
                <ArrowBack />
              </IconButton>
            </Box>
            <Typography variant="h5">{data.data.name}</Typography>
            {data.data.description && (
              <Typography>{data.data.description}</Typography>
            )}
          </Box>
        </Box>
      )}
    </Box>
  )
}

export default OrgDetail
