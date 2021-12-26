import { Box, Grid, Tab, Tabs, Typography } from '@mui/material'
import { makeStyles } from '@mui/styles'
import { NextPage } from 'next'
import { SyntheticEvent, useState } from 'react'
import { useCurrentUser } from '../../../hooks/useCurrentUser'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { PageProps } from '../../../types/page'
import { UserProfile } from '../../components/user/UserProfile'
import clsx from 'clsx'
import theme from '../../theme/theme'
import { Cli } from '../../components/user/Cli'

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

  const classes = useStyles(theme)

  return (
    <Box mt={2} width="100%">
      <Tabs value="1" onChange={handleChange} className={classes.tabs}>
        <Tab
          key="profile"
          value="profile"
          label="Profile"
          className={clsx(tab === 'profile' && classes.selected)}
          aria-selected={tab === 'profile'}
        />
        <Tab
          key="cli"
          value="cli"
          label="CLI"
          className={clsx(tab === 'cli' && classes.selected)}
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

const useStyles = makeStyles((theme) => ({
  tabs: {
    borderBottom: '1px solid black',
  },
  selected: {
    borderBottom: `2px solid ${theme.palette.primary.main}`,
  },
}))

export default UserPage
