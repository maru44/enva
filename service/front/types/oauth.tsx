export type cognitoCallbackParams = {
  code: string
  state: string
}

export type cognitoTokenResponse = {
  id_token: string
  access_token: string
  refresh_token: string
  expires_in: number
  token_type: string
  error: string
}
