import { NextPageContext } from 'next'
import { parseCookies } from 'nookies'
import { ApiUrl, IsDevelopment, ThisUrl } from '../config/env'

export const GetPath: { [key: string]: string } = {
  PROJECT_LIST_USER: '/project/list/user',
  PROJECT_LIST_ORG: '/project/list/org',
  PROJECT_DETAIl: '/project/detail',

  KVS_BY_PROJECT: '/kv',
  KV_CREATE: '/kv/create',
  KV_UPDATE: '/kv/update',
  KV_DELETE: '/kv/delete',
} as const
export type GetPath = typeof GetPath[keyof typeof GetPath]

type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE'

export const fetchBaseApi = async (
  path: string,
  method: HttpMethod,
  body?: { [key: string]: any },
  headers?: { [key: string]: string }
): Promise<Response> => {
  return fetch(`${ApiUrl}${path}`, {
    method: method,
    mode: 'cors',
    credentials: 'include',
    headers: headers ?? {
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: body && JSON.stringify(body),
  })
}

export async function fetcher(func: Promise<Response>): Promise<Response> {
  const res = await func
  switch (res.status) {
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

export async function fetcherGetFromApiUrl(path: string) {
  try {
    const fn = fetch(`${ApiUrl}${path}`, {
      method: 'GET',
      mode: 'cors',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json; charset=utf-8',
      },
    })

    const res = await fetcher(fn)
    const ret = await res.json()

    return ret
  } catch (e) {
    IsDevelopment && console.log(e)
    throw new Error('Internal Server Error')
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
