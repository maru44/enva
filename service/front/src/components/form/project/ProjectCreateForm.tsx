import { Box, Button, Grid, TextField, Typography, Paper } from '@mui/material'
import { useRouter } from 'next/router'
import { useSnackbar } from 'notistack'
import React, { useState } from 'react'
import { fetchCreateProject } from '../../../../http/project'
import { ProjectInput } from '../../../../types/project'
import { slugify } from '../../../../utils/slug'

export type ProjectCreateProps = {
  orgId?: string
}

export const ProjectCreateForm = ({ orgId }: ProjectCreateProps) => {
  const [slug, setSlug] = useState<string>('')
  const router = useRouter()
  const snack = useSnackbar()

  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const t = e.currentTarget
    const name = t.project_name.value
    const description = t.description.value === '' ? null : t.description.value

    const input: ProjectInput = {
      name: name,
      slug: slug,
      description: description,
    }
    const res = await fetchCreateProject(input)
    const ret = await res.json()
    if (res.status === 200) {
      const slug = ret['data']
      router.push(`/project/${slug}`)
    } else {
      const message = ret['error']
      snack.enqueueSnackbar(message, { variant: 'error' })
    }
  }

  const isPostable = (): boolean => {
    if (!slug) return false

    // @TODO user or orgs project

    return true
  }

  return (
    <Box width="100%" component="form" onSubmit={submit}>
      <Grid container>
        <Grid item xs={0} sm={1} md={3} />
        <Grid
          component={Paper}
          item
          xs={12}
          sm={10}
          md={6}
          display="flex"
          flexDirection="column"
          p={1}
          variant="outlined"
        >
          <Typography variant="h5">New Project</Typography>
          <Box mt={2}>
            <TextField
              name="project_name"
              variant="outlined"
              label="Name"
              required
              onChange={(e) => {
                setSlug(slugify(e.currentTarget.value))
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
            <Button type="submit" variant="outlined" disabled={!isPostable()}>
              Create
            </Button>
          </Box>
        </Grid>
      </Grid>
    </Box>
  )
}
