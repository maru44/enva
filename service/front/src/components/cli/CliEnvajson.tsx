import { Box, Grid, Typography } from '@mui/material'
import attributes from './attr.json'

type Attr = {
  attr: string
  required: boolean
  exp: string
}

export const CliEnvajson: React.FC = () => {
  const attrs: Attr[] = attributes ?? []

  return (
    <Box>
      <Box mt={6}>
        <Grid container>
          <Grid item xs={1} />
          <Grid item xs={10}>
            <Box pl={1} pr={1}>
              <Typography mb={4} variant="h5">
                enva.json
              </Typography>
              {attrs.map((a, i) => (
                <Box mb={2}>
                  <Typography mb={1} variant="h6" key={i}>
                    <b>{a.attr}</b> {a.required && '(required)'}
                  </Typography>
                  <Typography>
                    {a.exp}
                    {a.attr === 'file' && (
                      <Typography>
                        We support <b>.envrc</b>, <b>.yaml</b> (<b>.yml</b>) and{' '}
                        <b>.tfvars</b> extensions.
                        <br />
                        If you indicate extension other than those, it will be
                        written <b>key=value</b> format like <b>.env</b>.
                      </Typography>
                    )}
                  </Typography>
                </Box>
              ))}
            </Box>
          </Grid>
        </Grid>
      </Box>
    </Box>
  )
}
