import { Box, Grid, Tab, Tabs, Typography } from '@mui/material'
import { NextPage } from 'next'
import { SyntheticEvent, useState } from 'react'
import { useCurrentUser } from '../../../hooks/useCurrentUser'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { PageProps } from '../../../types/page'
import { UserProfile } from '../../components/user/UserProfile'
import clsx from 'clsx'
import { Cli } from '../../components/user/Cli'
import styles from '../../styles/user.module.css'

type tab = 'profile' | 'cli'

export type UserPageProps = {
  tabSelect?: tab
} & PageProps

const UserPage: NextPage<UserPageProps> = (props) => {
  useRequireLogin()
  const { currentUser } = useCurrentUser()
  const [tab, setTab] = useState<'profile' | 'cli'>(
    props.tabSelect ?? 'profile'
  )

  const handleChange = (e: SyntheticEvent, newValue: tab) => {
    setTab(newValue)
  }

  return (
    <Box mt={2} width="100%">
      <Tabs value="1" onChange={handleChange} className={styles.tabs}>
        <Tab
          key="profile"
          value="profile"
          label="Profile"
          className={clsx(tab === 'profile' && styles.selected)}
          aria-selected={tab === 'profile'}
        />
        <Tab
          key="cli"
          value="cli"
          label="CLI"
          className={clsx(tab === 'cli' && styles.selected)}
          aria-selected={tab === 'cli'}
        />
      </Tabs>
      <Grid mt={1}>
        {currentUser && tab === 'profile' && (
          <UserProfile currentUser={currentUser} />
        )}
        {currentUser && tab === 'cli' && <Cli />}
      </Grid>
    </Box>
  )
}

export default UserPage
