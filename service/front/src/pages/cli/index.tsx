import { Box, Tab, Tabs, Typography } from '@mui/material'
import { NextPage } from 'next'
import { PageProps } from '../../../types/page'
import { SyntheticEvent, useEffect, useState } from 'react'
import styles from '../../styles/cli.module.css'
import { useRouter } from 'next/router'
import { CliInstallations } from '../../components/cli/CliInstallations'
import { CliCommands } from '../../components/cli/CliCommands'

type tab = 'installations' | 'commands'

const CliIndex: NextPage<PageProps> = (props) => {
  const router = useRouter()
  const [tab, setTab] = useState<tab>('installations')
  const handleChange = (e: SyntheticEvent, newValue: tab) => {
    setTab(newValue)
  }

  const { page } = router.query
  useEffect(() => {
    page === 'commands' && setTab('commands')
  }, [page])

  return (
    <Box>
      <Box mt={6}>
        <Typography variant="h4">CLI</Typography>
      </Box>
      <Box mt={6}>
        <Tabs value={tab} onChange={handleChange} className={styles.tabs}>
          <Tab
            key="installations"
            value="installations"
            label="Installations"
          />
          <Tab key="commands" value="commands" label="Commands" />
        </Tabs>
        {tab === 'installations' && <CliInstallations />}
        {tab === 'commands' && <CliCommands />}
      </Box>
    </Box>
  )
}

export default CliIndex
