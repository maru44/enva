import { FileCopy } from '@material-ui/icons'
import {
  Box,
  Button,
  Grid,
  IconButton,
  TextField,
  Paper,
  Tooltip,
  Typography,
} from '@mui/material'
import { useSnackbar } from 'notistack'
import { useState } from 'react'
import useSWR, { mutate } from 'swr'
import { useRequireLogin } from '../../../hooks/useRequireLogin'
import { cliUserResponseBody } from '../../../http/body/cliUser'
import { fetchUpdateCliUser } from '../../../http/cliUser'
import { fetcherGetFromApiUrl, GetPath } from '../../../http/fetcher'

export const Cli: React.FC = () => {
  const [pass, setPass] = useState<string | undefined>(undefined)
  const [isCopied, setIsCopied] = useState<boolean>(false)
  const snack = useSnackbar()

  const { data, error } = useSWR<cliUserResponseBody>(
    GetPath.CLI_USER,
    fetcherGetFromApiUrl
  )

  const gen = async () => {
    try {
      const res = await fetchUpdateCliUser()
      const ret: cliUserResponseBody = await res.json()

      switch (res.status) {
        case 200:
          setPass(ret.data)
          mutate(GetPath.CLI_USER)
          return
        default:
          snack.enqueueSnackbar(ret.error, { variant: 'error' })
          return
      }
    } catch {
      snack.enqueueSnackbar('Internal Server Error', { variant: 'error' })
    }
  }

  useRequireLogin()

  return (
    <Grid container mt={2}>
      <Grid xs={0} sm={2} md={3} />
      <Grid
        item
        xs={12}
        sm={8}
        md={6}
        component={Paper}
        pt={2}
        pb={2}
        pl={1}
        pr={1}
        variant="outlined"
      >
        <Typography variant="h5">Password For CLI</Typography>
        <Box mt={2}>
          <Typography>This password can be used in only cli.</Typography>
          <br />
          {data && data.data && (
            <Typography>
              Your password has been already generated.
              <br />
              If you forgot it, regenerate from here.
            </Typography>
          )}
          {data && !data.data && (
            <Typography>
              If you want to use cli, you must generate password for cli.
            </Typography>
          )}
        </Box>
        {pass && (
          <Box mt={4}>
            <Box mb={1}>
              <Typography variant="h6">new password</Typography>
              <Typography>Keep it secret.</Typography>
            </Box>
            <TextField
              variant="outlined"
              type="text"
              value={pass}
              fullWidth
              multiline
              InputProps={{
                endAdornment: (
                  <Box ml={1}>
                    <Tooltip
                      title="copied!"
                      disableHoverListener
                      open={isCopied}
                      placement="top"
                      arrow
                    >
                      <IconButton
                        onClick={() => {
                          navigator.clipboard.writeText(pass)
                          setIsCopied(true)
                        }}
                      >
                        <FileCopy />
                      </IconButton>
                    </Tooltip>
                  </Box>
                ),
              }}
            />
          </Box>
        )}
        <Box mt={2} textAlign="right">
          {data && data.data ? (
            <Button
              type="button"
              variant="contained"
              onClick={() => {
                gen()
                setIsCopied(false)
              }}
            >
              Re Generate
            </Button>
          ) : (
            <Button type="button" variant="contained" onClick={() => gen()}>
              Generate
            </Button>
          )}
        </Box>
      </Grid>
    </Grid>
  )
}
