import { FileCopy } from '@material-ui/icons'
import {
  Box,
  Button,
  Grid,
  IconButton,
  TextField,
  Tooltip,
  Typography,
} from '@mui/material'
import { NextPage } from 'next'
import { useState } from 'react'
import useSWR from 'swr'
import { cliUserResponseBody } from '../../../../http/body/cliUser'
import {
  fetchCreateCliUser,
  fetchUpdateCliUser,
} from '../../../../http/cliUser'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'
import { PageProps } from '../../../../types/page'

const CliPassword: NextPage<PageProps> = (props) => {
  const [pass, setPass] = useState<string | undefined>(undefined)
  const [isCopied, setIsCopied] = useState<boolean>(false)

  const { data, error } = useSWR<cliUserResponseBody>(
    GetPath.CLI_USER,
    fetcherGetFromApiUrl
  )

  const gen = async (re: boolean) => {
    try {
      const res = re ? await fetchUpdateCliUser() : await fetchCreateCliUser()
      const ret: cliUserResponseBody = await res.json()

      switch (res.status) {
        case 200:
          setPass(ret.data)
          break
        default:
          break
      }
    } catch (e) {
      console.log(e)
    }
  }

  return (
    <Grid container mt={2}>
      <Grid item xs={12}>
        <Typography variant="h5">Password For CLI</Typography>
      </Grid>
      <Grid item mt={2} xs={12}>
        <Typography>This password can be used in only cli.</Typography>
      </Grid>
      <Grid item mt={4} xs={12}>
        <Grid sm={6} md={4}>
          {data && data.data && (
            <Box>
              <Typography>
                Your password has been already generated.
                <br />
                If you forgot it. {`>>`}
              </Typography>
            </Box>
          )}
          {data && !data.data && (
            <Typography>
              If you want to use cli, you must generate password for cli.
            </Typography>
          )}
        </Grid>
      </Grid>
      {pass && (
        <Grid item mt={4} xs={12}>
          <Grid sm={6} md={4}>
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
          </Grid>
        </Grid>
      )}
      <Grid mt={2} item xs={12} sm={6} md={4} textAlign="right">
        {data && data.data ? (
          <Button
            type="button"
            variant="contained"
            onClick={() => {
              gen(true)
              setIsCopied(false)
            }}
          >
            Re Generate
          </Button>
        ) : (
          <Button type="button" variant="contained" onClick={() => gen(false)}>
            Generate
          </Button>
        )}
      </Grid>
    </Grid>
  )
}

export default CliPassword
