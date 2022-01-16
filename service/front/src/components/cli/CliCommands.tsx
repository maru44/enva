import { Box, Grid, Typography } from '@mui/material'
import commands from './explain.json'

type Command = {
  command: string
  explain: string
}

export const CliCommands: React.FC = () => {
  const cmds: Command[] = commands ?? []

  return (
    <Box>
      <Grid container>
        <Grid item xs={1} />
        <Grid item xs={10}>
          <Box mt={6} pl={1} pr={1}>
            <Typography variant="h5">List of CLI Commands</Typography>
          </Box>
          <Box mt={4} pl={1} pr={1}>
            {cmds &&
              cmds.length > 0 &&
              cmds.map((c, i) => (
                <Box mb={2} key={i}>
                  <Typography variant="h6">{c.command}</Typography>
                  <Typography mt={1}>
                    <pre>{deleteTab(c.explain)}</pre>
                  </Typography>
                </Box>
              ))}
          </Box>
        </Grid>
      </Grid>
    </Box>
  )
}

const deleteTab = (raw: string) => raw.replaceAll('\t', '')
