import {
//   GetServerSidePropsContext,
//   GetStaticPathsContext,
//   GetStaticPropsContext,
  NextPageContext,
} from 'next'
import { parseCookies } from 'nookies'
import { ThisUrl } from '../config/env'

export async function fetcher(func: Promise<Response>) {
  const res = await func
  switch (res.status) {
    case 200:
      return res
    case 401:
      await fetch(`${ThisUrl}/api/auth/refresh`, {
        method: 'GET',
        credentials: 'include',
      })
      const res2 = await func
      return res2
    default:
      return res
  }
}

export async function fetcherSSR<
  F extends (...args: any[]) => Promise<Response>
>(ctx: NextPageContext, func: F) {
  const cookies = parseCookies(ctx)
  const res = await func(cookies.accessToken ?? '')
  switch (res.status) {
    case 200:
      return res
    case 401:
      const resUser = await fetch(`${ThisUrl}/api/auth/server/refresh`, {
        method: 'GET',
        credentials: 'include',
      })
      // setCookie
      if (resUser.status === 200) {
        // destroyCookie(ctx, 'accessToken');
        // destroyCookie(ctx, 'refreshToken');
        // const data = await resUser.json();
        // mySetCookie('refreshToken', data['refreshToken'], 10 * 60 * 60, ctx);
        // const newAccessToken = data['idToken'];
        // mySetCookie('accessToken', newAccessToken, 5 * 60, ctx);
        // const res2 = await func;
        // return res2;
      } else {
        return res
      }
    default:
      return res
  }
}
