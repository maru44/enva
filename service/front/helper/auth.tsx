import { callbackUrl } from '../config/aws'

export const getCognitoToken = async (code: string, state: string) => {
  return await fetch(
    `https://${process.env.NEXT_PUBLIC_COGNITO_DOMAINNAME}.auth.${
      process.env.NEXT_PUBLIC_COGNITO_REGION
    }.amazoncognito.com/oauth2/token?grant_type=authorization_code&code=${code}&state=${state}&client_id=${
      process.env.COGNITO_CLIENT_ID
    }&client_secret=${
      process.env.COGNITO_CLIENT_SECRET
    }&redirect_uri=${encodeURI(callbackUrl)}`,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    }
  )
}

export const refreshCognitoToken = async (refreshToken: string) => {
  return await fetch(
    `https://${process.env.NEXT_PUBLIC_COGNITO_DOMAINNAME}.auth.${process.env.NEXT_PUBLIC_COGNITO_REGION}.amazoncognito.com/oauth2/token?grant_type=refresh_token&client_id=${process.env.COGNITO_CLIENT_ID}&client_secret=${process.env.COGNITO_CLIENT_SECRET}&refresh_token=${refreshToken}`,
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
    }
  )
}
