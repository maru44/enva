import { Box, Grid, Typography } from '@mui/material'
import commands from './explain.json'
import styles from '../../styles/cli.module.css'

type Command = {
  command: string
  explain: string
}

export const CliCommands: React.FC = () => {
  const cmds: Command[] = commands ?? []

  return (
    <Box>
      <Box mt={6}>
        <Grid container>
          <Grid item xs={1} />
          <Grid item xs={10}>
            <Box pl={1} pr={1}>
              <Typography variant="h5">enva.json</Typography>
              <Typography variant="subtitle1" mt={4}>
                We support <b>.envrc</b>, <b>.yaml</b> (<b>.yml</b>) and{' '}
                <b>.tfvars</b> extensions.
                <br />
                If you indicate extension other than those, it will be written{' '}
                <b>key=value</b> format like <b>.env</b>.
              </Typography>
            </Box>
          </Grid>
        </Grid>
      </Box>
      <Grid container>
        <Grid item xs={1} />
        <Grid item xs={10}>
          <Box mt={8} pl={1} pr={1}>
            <Typography variant="h5">List of CLI Commands</Typography>
          </Box>
          <Box mt={4} pl={1} pr={1}>
            {cmds &&
              cmds.length > 0 &&
              cmds.map((c, i) => (
                <Box mb={4} key={i}>
                  <Typography variant="h6" mb={1}>
                    <b>{c.command}</b>
                  </Typography>
                  <Typography className={styles.preWrap} mt={1}>
                    {deleteTab(c.explain)}
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
