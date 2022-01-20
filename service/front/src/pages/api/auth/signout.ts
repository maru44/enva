import { NextApiRequest, NextApiResponse } from 'next'
import { serialize } from 'cookie'
import {
  CookieKeyAccessToken,
  CookieKeyIdToken,
  CookieKeyRefreshToken,
  getCookieOption,
} from '../../../../config/cookie'

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  try {
    res.setHeader('Set-Cookie', [
      serialize(CookieKeyIdToken, '', getCookieOption(0)),
      serialize(CookieKeyAccessToken, '', getCookieOption(0)),
      serialize(CookieKeyRefreshToken, '', getCookieOption(0)),
    ])
    res.status(200).json({ message: 'succeeded to sign out' })
  } catch {
    res.status(500).json({ message: 'Internal Server Error' })
  }
  return
}
