import { Box, Button, TextField } from '@mui/material'
import { makeStyles } from '@mui/styles'
import clsx from 'clsx'
import React, { useState } from 'react'
import { fetcher } from '../../../../http/fetcher'
import { fetchCreateProject } from '../../../../http/project'
import { ProjectInput } from '../../../../types/project'
import { slugify } from '../../../../utils/slug'

export type ProjectCreateProps = {
  orgId?: string
}

export const ProjectCreateForm = ({ orgId }: ProjectCreateProps) => {
  const [slug, setSlug] = useState<string>('')

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
    const res = await fetcher(fetchCreateProject(input))
    const ret = await res.json()
    if (res.status === 200) {
      const id = ret['data']
      console.log(id) // @TODO fix
    } else {
      const message = ret['error']
      console.log(message) // @TODO fix
    }
  }

  const classes = useStyle()

  const isPostable = (): boolean => {
    if (!slug) return false

    // @TODO user or orgs project

    return true
  }

  return (
    <Box>
      <form onSubmit={submit}>
        <Box display="flex" flexDirection="column">
          <TextField
            name="project_name"
            variant="outlined"
            label="Name"
            required
            onChange={(e) => {
              setSlug(slugify(e.currentTarget.value))
            }}
            className={clsx(classes.textField)}
          />
          <TextField
            name="description"
            label="Description"
            multiline
            rows={6}
            className={clsx(classes.textField)}
          />
          <Button type="submit" variant="outlined" disabled={!isPostable()}>
            Create
          </Button>
        </Box>
      </form>
    </Box>
  )
}

const useStyle = makeStyles(() => ({
  textField: {
    marginBottom: 8,
  },
}))
