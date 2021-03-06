import {
  Box,
  Grid,
  IconButton,
  Tab,
  Tabs,
  Tooltip,
  Typography,
} from '@mui/material'
import Versions from '../../../public/enva/tar.json'
import { SyntheticEvent, useState } from 'react'
import Link from 'next/link'
import { CloudDownload, FileCopy } from '@material-ui/icons'
import styles from '../../styles/cli.module.css'
import { tarJson } from '../../../types/os'

type tab = 'linux' | 'darwin' | 'windows'

export const CliInstallations: React.FC = (props) => {
  // sort desc
  const vs: tarJson[] = Versions
    ? Versions.sort((a, b) => b.version.localeCompare(a.version))
    : []

  const [tab, setTab] = useState<tab | 'history'>('linux')
  const handleChange = (e: SyntheticEvent, newValue: tab) => {
    setTab(newValue)
  }

  return (
    <Box>
      <Grid container mt={6}>
        <Grid item xs={1} />
        <Grid item xs={10}>
          <Box>
            <Typography variant="h5">CLI Binaries</Typography>
          </Box>
          <Box mt={2}>
            <Box mt={2}>
              <Typography>
                A sample code to install enva CLI command.
              </Typography>
            </Box>
            <Box mt={2} p={2} className={styles.code}>
              <Typography
                alignItems="center"
                justifyContent="space-between"
                display="flex"
              >
                <code>
                  rm -f /usr/local/bin/enva {`&&`} tar -C /usr/local/bin -xvzf
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
          {vs && vs.length > 0 && (
            <Box mt={6}>
              <Tabs value={tab} onChange={handleChange} className={styles.tabs}>
                <Tab key="linux" value="linux" label="Linux" />
                <Tab key="darwin" value="darwin" label="Mac" />
                <Tab key="windows" value="windows" label="Windows" />
                <Tab key="history" value="history" label="history" />
              </Tabs>
              {tab !== 'history' && (
                <Box mt={6}>
                  <TarHrefs version={vs[0]} tab={tab} />
                </Box>
              )}
              {tab === 'history' && (
                <Box mt={6}>
                  {vs.map((v, i) => (
                    <Box key={i} mb={2}>
                      <Box display="flex" alignItems="center">
                        <Typography variant="h6">{v.version}</Typography>
                        <Typography variant="h6" ml={2}>
                          {v.updated_at}
                        </Typography>
                      </Box>
                      <TarHrefs version={v} />
                    </Box>
                  ))}
                </Box>
              )}
            </Box>
          )}
        </Grid>
      </Grid>
    </Box>
  )
}

const TarHrefs: React.FC<{ version: tarJson; tab?: tab }> = ({
  version,
  tab,
}) => {
  return (
    <Box>
      {version.oss &&
        version.oss.map((os, i) => (
          <Box key={i}>
            {(!tab || os.os === tab) &&
              os.archs &&
              os.archs.map((a, i) => (
                <TarHref
                  key={i}
                  fileName={fileName(os.os, a, version.version)}
                />
              ))}
          </Box>
        ))}
    </Box>
  )
}

const TarHref: React.FC<{ fileName: string }> = ({ fileName }) => {
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

const fileName = (os: string, arch: string, version: string) =>
  `enva_${version}_${os}_${arch}.tar.gz`
