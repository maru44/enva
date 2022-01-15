import { IsDevelopment } from './env'
import { randomSlug } from '../utils/random'

const awsConfig = {
  identityPoolId: process.env.NEXT_PUBLIC_COGNITO_IDENTITYPOOLID,
  region: process.env.NEXT_PUBLIC_COGNITO_REGION,
  UserPoolId: process.env.NEXT_PUBLIC_COGNITO_USERPOOLID,
  userPoolWebClientId: process.env.NEXT_PUBLIC_COGNITO_WEBCLIENTID,
}
export default awsConfig

const cognitoUrl = `https://${process.env.NEXT_PUBLIC_COGNITO_DOMAINNAME}.auth.${process.env.NEXT_PUBLIC_COGNITO_REGION}.amazoncognito.com/`

export const callbackUrl = IsDevelopment
  ? 'http://localhost:3000/api/auth/callback/cognito'
  : ''

export const loginUrl = `${cognitoUrl}login?response_type=${
  process.env.NEXT_PUBLIC_COGNITO_RESPONSE_TYPE
}&client_id=${
  process.env.NEXT_PUBLIC_COGNITO_USERPOOLWEBCLIENTID
}&state=${randomSlug(
  12
)}&scope=openid%20email%20profile&redirect_uri=${encodeURI(callbackUrl)}`
