import { Box, Button, TextField, Tooltip, Typography } from '@mui/material'
import { Auth } from 'aws-amplify'
import { NextPage } from 'next'
import React, { useReducer } from 'react'
import {
  initialSignUpState,
  signUpReducer,
  signUpState,
} from '../../../states/cognito/signUpReducer'
import { PageProps } from '../../../types/page'

const SignUpPage: NextPage<PageProps> = (props) => {
  const [state, dispatch] = useReducer(signUpReducer, initialSignUpState)
  const signUp = async () => {
    try {
      const email = state.email!
      const password = state.password!
      const username = state.username!
      await Auth.signUp({
        username,
        password,
        attributes: {
          email,
        },
      })
    } catch (e) {
      console.log(e)
    }
  }

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    signUp()
  }

  return (
    <Box>
      <Box mt={6}>
        <Box>
          <Typography variant="h4">Sign up</Typography>
        </Box>
        <Box mt={4} component="form" onSubmit={onSubmit}>
          <Box>
            <Tooltip title="4+ length is required" arrow placement="top">
              <TextField
                onChange={(e) =>
                  dispatch({
                    type: 'setUsername',
                    value: e.currentTarget.value,
                  })
                }
                variant="outlined"
                required
                fullWidth
                label="username"
              />
            </Tooltip>
          </Box>
          <Box mt={4}>
            <TextField
              type="email"
              onChange={(e) =>
                dispatch({ type: 'setEmail', value: e.currentTarget.value })
              }
              variant="outlined"
              required
              fullWidth
              label="email"
            />
          </Box>
          <Box mt={4}>
            <TextField
              onChange={(e) =>
                dispatch({ type: 'setPassword', value: e.currentTarget.value })
              }
              variant="outlined"
              required
              fullWidth
              label="password"
              type="password"
            />
          </Box>
          <Box mt={4}>
            <TextField
              type="password"
              onChange={(e) =>
                dispatch({ type: 'setPassword2', value: e.currentTarget.value })
              }
              variant="outlined"
              required
              fullWidth
              label="password confirm"
            />
          </Box>
          <Box mt={8} textAlign="right">
            <Button
              type="submit"
              variant="contained"
              disabled={!state.canSubmit}
            >
              Sign up
            </Button>
          </Box>
        </Box>
      </Box>
    </Box>
  )
}

export default SignUpPage
