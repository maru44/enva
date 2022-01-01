import { Box, Button, Grid, Paper, TextField, Typography } from '@mui/material'
import { useRouter } from 'next/router'
import { useSnackbar } from 'notistack'
import { ReactEventHandler, useState } from 'react'
import { OrgCreateResponseBody } from '../../../../http/body/org'
import { fetchCreateOrg } from '../../../../http/org'
import { OrgInput } from '../../../../types/org'
import { slugify } from '../../../../utils/slug'

export const OrgCreateForm = () => {
  const [input, setInput] = useState<OrgInput | undefined>(undefined)
  const router = useRouter()
  const snack = useSnackbar()

  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (!input) return

    const res = await fetchCreateOrg(input)
    const ret: OrgCreateResponseBody = await res.json()
    if (res.status === 200) {
      const slug = ret.data
      router.push(`/org/${slug}`)
    } else {
      const message = ret.error
      snack.enqueueSnackbar(message, { variant: 'error' })
    }
  }

  return (
    <Box width="100%" component="form" onSubmit={submit}>
      <Grid container>
        <Grid item xs={0} sm={1} md={3} />
        <Grid
          item
          component={Paper}
          xs={12}
          sm={10}
          md={6}
          display="flex"
          flexDirection="column"
          p={1}
          variant="outlined"
        >
          <Typography variant="h5">New Organization</Typography>
          <Box mt={2}>
            <TextField
              name="org_name"
              variant="outlined"
              label="name"
              required
              onChange={(e) => {
                setInput({
                  slug: slugify(e.currentTarget.value),
                  name: e.currentTarget.value,
                })
              }}
              fullWidth
              inputProps={{ maxLength: 32 }}
            />
          </Box>
          <Box mt={2}>
            <TextField
              name="description"
              label="Description"
              multiline
              rows={6}
              fullWidth
            />
          </Box>
          <Box mt={2} textAlign="right">
            <Button type="submit" variant="outlined" disabled={!input?.slug}>
              Create
            </Button>
          </Box>
        </Grid>
      </Grid>
    </Box>
  )
}
