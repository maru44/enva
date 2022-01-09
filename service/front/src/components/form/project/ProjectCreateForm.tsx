import { AccountCircle, Apartment } from '@material-ui/icons'
import {
  Box,
  Button,
  Grid,
  TextField,
  Typography,
  Paper,
  Select,
  MenuItem,
  IconButton,
} from '@mui/material'
import { useRouter } from 'next/router'
import { useSnackbar } from 'notistack'
import React, { useState } from 'react'
import useSWR from 'swr'
import { OrgsResponseBody } from '../../../../http/body/org'
import { projectCreateResponseBody } from '../../../../http/body/project'
import { fetcherGetFromApiUrl, GetPath } from '../../../../http/fetcher'
import { fetchCreateProject } from '../../../../http/project'
import { ProjectInput } from '../../../../types/project'
import { slugify } from '../../../../utils/slug'

export type ProjectCreateProps = {
  orgId?: string
}

export const ProjectCreateForm: React.FC<ProjectCreateProps> = ({ orgId }) => {
  const [slug, setSlug] = useState<string>('')
  const [orgSlug, setOrgSlug] = useState<string | undefined>(undefined)
  const router = useRouter()
  const snack = useSnackbar()

  const { data, error } = useSWR<OrgsResponseBody, ErrorConstructor>(
    GetPath.ORG_ADMIN_LIST,
    fetcherGetFromApiUrl
  )

  const submit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const t = e.currentTarget
    const orgId = t.orgs.value === 'user' ? null : t.orgs.value
    const name = t.project_name.value
    const description = t.description.value === '' ? null : t.description.value

    const input: ProjectInput = {
      name: name,
      slug: slug,
      description: description,
      org_id: orgId,
    }
    const res = await fetchCreateProject(input)
    const ret: projectCreateResponseBody = await res.json()
    if (res.status === 200) {
      const slug = ret.data
      const path = input.org_id
        ? `/project/${orgSlug}/${slug}/`
        : `/project/${slug}`
      router.push(path)
    } else {
      const message = ret.error
      snack.enqueueSnackbar(message, { variant: 'error' })
    }
  }

  const isPostable = (): boolean => {
    if (!slug) return false

    // @TODO user or orgs project

    return true
  }

  return (
    <Box
      width="100%"
      component="form"
      onSubmit={(e: React.FormEvent<HTMLFormElement>) => {
        submit(e)
      }}
    >
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
            <Select label="orgs" name="orgs" fullWidth defaultValue="user">
              <MenuItem value="user">
                <Box display="flex" flexDirection="row" alignItems="center">
                  <IconButton>
                    <AccountCircle />
                  </IconButton>
                  <Typography>As a user</Typography>
                </Box>
              </MenuItem>
              {data &&
                data.data &&
                data.data.map((o, i) => (
                  <MenuItem
                    key={i}
                    value={o.id}
                    onClick={() => {
                      setOrgSlug(o.slug)
                    }}
                  >
                    <Box display="flex" flexDirection="row" alignItems="center">
                      <IconButton>
                        <Apartment />
                      </IconButton>
                      <Typography>{o.name}</Typography>
                    </Box>
                  </MenuItem>
                ))}
            </Select>
          </Box>
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
