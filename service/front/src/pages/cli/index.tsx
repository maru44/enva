import { Box, IconButton, Tab, Tabs, Tooltip, Typography } from '@mui/material'
import { NextPage } from 'next'
import { PageProps } from '../../../types/page'
import Versions from '../../../public/enva/version.json'
import { OsArch } from '../../../types/os'
import { SyntheticEvent, useState } from 'react'
import Link from 'next/link'
import { CloudDownload, FileCopy } from '@material-ui/icons'
import styles from '../../styles/cli.module.css'

type tab = 'linux' | 'mac' | 'windows'

type props = {
  version: string
}

const osArch: { [key: string]: string[] } = {
  mac: ['amd64'],
  linux: ['386', 'amd64'],
  windows: ['386', 'amd64'],
}

const CliIndex: NextPage<PageProps> = (props) => {
  // sort desc
  const vs = Versions ? Versions.sort((a, b) => b.localeCompare(a)) : []

  const [tab, setTab] = useState<tab | 'history'>('linux')
  const [isCopied, setIsCopied] = useState<boolean>(false)

  const handleChange = (e: SyntheticEvent, newValue: tab) => {
    setTab(newValue)
  }

  return (
    <Box>
      <Box mt={6}>
        <Typography variant="h5">Cli Binaries</Typography>
      </Box>
      <Box mt={4}>
        <Typography variant="h6">Installation</Typography>
        <Box mt={2}>
          <Box mt={2}>
            <Typography>A sample code to install enva cli command.</Typography>
          </Box>
          <Box mt={2} p={2} className={styles.code}>
            <Typography
              alignItems="center"
              justifyContent="space-between"
              display="flex"
            >
              <code>
                rm -f /usr/local/bin/enva -C /usr/local/bin -xvzf
                enva_v1.0.0_linux_amd64.tar.gz
              </code>
              <Tooltip title="copy" placement="top" arrow>
                <IconButton
                  onClick={() => {
                    navigator.clipboard.writeText(
                      'rm -f /usr/local/bin/enva -C /usr/local/bin -xvzf enva_v1.0.0_linux_amd64.tar.gz'
                    )
                  }}
                >
                  <FileCopy />
                </IconButton>
              </Tooltip>
            </Typography>
          </Box>
        </Box>
      </Box>
      {vs && vs.length > 0 && (
        <Box mt={6}>
          <Tabs value={tab} onChange={handleChange} className={styles.tabs}>
            <Tab key="linux" value="linux" label="Linux" />
            <Tab key="mac" value="mac" label="Mac" />
            <Tab key="windows" value="windows" label="Windows" />
            <Tab key="history" value="history" label="history" />
          </Tabs>
          {tab !== 'history' && (
            <Box mt={6}>
              <Box>
                {osArch[tab].map((a, i) => (
                  <Box key={i}>
                    <TarHref fileName={fileName(tab, a, vs[0])} />
                  </Box>
                ))}
              </Box>
            </Box>
          )}
          {tab === 'history' && (
            <Box mt={6}>
              {vs.map((v, i) => (
                <Box key={i} mb={2}>
                  <Box>
                    <Typography variant="h6">{v}</Typography>
                  </Box>
                  {OsArch.map((oa, ii) => (
                    <TarHref key={ii} fileName={`enva_${v}_${oa}.tar.gz`} />
                  ))}
                </Box>
              ))}
            </Box>
          )}
        </Box>
      )}
      {}
    </Box>
  )
}

type tarProps = {
  fileName: string
}

const TarHref: React.FC<tarProps> = ({ fileName }) => {
  return (
    <Box display="flex" alignItems="center">
      <Typography>{fileName}</Typography>
      <Box ml={6}>
        <Link href={`/enva/${fileName}`} passHref>
          <IconButton>
            <CloudDownload />
          </IconButton>
        </Link>
      </Box>
    </Box>
  )
}

const fileName = (os: tab, arch: string, version: string) =>
  `enva_${version}_${os}_${arch}.tar.gz`

export default CliIndex
