import { Box, Grid, Tab, Tabs, Typography } from '@mui/material'
import { NextPage } from 'next'
import { SyntheticEvent, useState } from 'react'
import { useCurrentUser } from '../../../hooks/useCurrentUser'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { PageProps } from '../../../types/page'
import { UserProfile } from '../../components/user/UserProfile'
import { Cli } from '../../components/user/Cli'
import styles from '../../styles/user.module.css'
import { UserOrgs } from '../../components/user/UserOrgs'
import { SignOut } from '../../components/user/SignOut'

type tab = 'profile' | 'cli' | 'orgs' | 'out'

export type UserPageProps = {
  tabSelect?: tab
} & PageProps

const UserPage: NextPage<UserPageProps> = (props) => {
  useRequireLogin()
  const { currentUser } = useCurrentUser()
  const [tab, setTab] = useState<tab>(props.tabSelect ?? 'profile')

  const handleChange = (e: SyntheticEvent, newValue: tab) => {
    setTab(newValue)
  }

  return (
    <Box mt={6} width="100%">
      <Tabs value={tab} onChange={handleChange} className={styles.tabs}>
        <Tab key="profile" value="profile" label="Profile" />
        <Tab key="cli" value="cli" label="CLI" />
        <Tab key="orgs" value="orgs" label="Orgs" />
        <Tab key="out" value="out" label="Sign out" />
      </Tabs>
      <Box mt={6}>
        {currentUser && tab === 'profile' && (
          <UserProfile currentUser={currentUser} />
        )}
        {currentUser && tab === 'cli' && <Cli />}
        {currentUser && tab === 'orgs' && <UserOrgs />}
        {currentUser && tab === 'out' && <SignOut />}
      </Box>
    </Box>
  )
}

export default UserPage
