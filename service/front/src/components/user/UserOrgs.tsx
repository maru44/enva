import { Box, Grid, IconButton } from '@mui/material'
import useSWR from 'swr'
import { OrgsResponseBody } from '../../../http/body/org'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'
import { CommonListCard } from '../CommonListCard'
import styles from '../../styles/org.module.css'
import Link from 'next/link'
import { AddCircle, Apartment } from '@material-ui/icons'

export const UserOrgs: React.FC = () => {
  const { data, error } = useSWR<OrgsResponseBody>(
    GetPath.ORG_LIST,
    fetcherGetFromApiUrl
  )

  //   @TODO if empty
  if (data && data.data && data.data.length === 0) {
    return <Box></Box>
  }

  return (
    <Grid container>
      <Box width="100%" alignItems="center" display="flex" justifyContent="end">
        <Link href="/org/create" passHref>
          <Box>
            <IconButton>
              <AddCircle />
            </IconButton>
            Create Org
          </Box>
        </Link>
      </Box>
      <Grid container mt={1} rowSpacing={2} columnSpacing={2}>
        {data &&
          data.data &&
          data.data.map((o, i) => (
            <CommonListCard
              key={i}
              info={o}
              linkAs={`/org/${o.slug}`}
              linkHref="/org/[slug]"
              icon={<Apartment />}
              styles={styles}
            />
          ))}
      </Grid>
    </Grid>
  )
}
