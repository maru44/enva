import { CookieSerializeOptions } from 'cookie'
import { IsProduction } from './env'

export const CookieKeyIdToken = 'id_token'
export const CookieKeyAccessToken = 'access_token'
export const CookieKeyRefreshToken = 'refresh_token'

export const getCookieOption = (age?: number): CookieSerializeOptions => {
  return {
    path: '/',
    domain: `${process.env.NEXT_PUBLIC_DOMAIN}`,
    httpOnly: true,
    secure: IsProduction,
    sameSite: IsProduction ? 'none' : 'lax',
    maxAge: age,
  }
}
