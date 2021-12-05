import { NextApiRequest, NextApiResponse } from 'next'
import { serialize } from 'cookie'
import { refreshCognitoToken } from '../../../../helper/auth'
import {
  CookieKeyAccessToken,
  CookieKeyIdToken,
  CookieKeyRefreshToken,
  getCookieOption,
} from '../../../../config/cookie'
import { cognitoTokenResponse } from '../../../../types/oauth'

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  const refreshToken = req.cookies.refresh_token

  try {
    // request to cognito token endpoint
    const response = await refreshCognitoToken(refreshToken)
    const ret: cognitoTokenResponse = await response.json()

    switch (response.status) {
      case 200:
        // set cookie
        res.setHeader('Set-Cookie', [
          serialize(
            CookieKeyIdToken,
            ret.id_token,
            getCookieOption(ret.expires_in)
          ),
          serialize(
            CookieKeyAccessToken,
            ret.access_token,
            getCookieOption(ret.expires_in)
          ),
          serialize(
            CookieKeyRefreshToken,
            ret.refresh_token ?? refreshToken,
            getCookieOption(3600 * 24 * 7 * 3)
          ),
        ])
        res.status(200).json({ message: 'succeeded to refresh token' })
        return
      default:
        res.status(400).json(ret)
        return
    }
  } catch (e) {
    res.status(500).json({ message: 'Internal Server Error' })
    return
  }
}
