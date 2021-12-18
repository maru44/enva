import { Box, Button, Input, Typography } from '@mui/material'
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
    <Box>
      <Typography variant="h2">Password For CLI</Typography>
      <Box m={4}>
        <Box>
          {data && data.data && (
            <p>Your password has been already generated.</p>
          )}
          {data && !data.data && (
            <p>Your must generate password for cli.</p>
          )}
        </Box>
        <Box>{pass && <Input type="text" value={pass} />}</Box>
        {data && data.data ? (
          <Button type="button" onClick={() => gen(true)}>
            Re Generate
          </Button>
        ) : (
          <Button type="button" onClick={() => gen(false)}>
            Generate
          </Button>
        )}
      </Box>
    </Box>
  )
}

export default CliPassword
