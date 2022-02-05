import { NextPageContext } from 'next'
import { parseCookies } from 'nookies'
import { IsDevelopment } from '../config/env'

export const GetPath = {
  PROJECT_LIST: '/project',
  PROJECT_LIST_USER: '/project/list/user',
  PROJECT_LIST_ORG: '/project/list/org', // ?id=
  PROJECT_DETAIl: '/project/detail',
  PROJECT_CREATE: '/project/create',
  PROJECT_DELETE: '/project/delete',
  PROJECT_UPDATE: '/project/update',

  KVS_BY_PROJECT: '/kv',
  KV_CREATE: '/kv/create',
  KV_UPDATE: '/kv/update',
  KV_DELETE: '/kv/delete',

  ORG_CREATE: '/org/create',
  ORG_LIST: '/org',
  ORG_ADMIN_LIST: '/org/admins',
  ORG_DETAIL: '/org/detail',

  ORG_INVITE: '/invite',
  ORG_INVITATION_DENY: '/invite/deny',
  ORG_INVITATION_DETAIL: '/invite/detail',
  ORG_INVITATION_ACCEPT: '/member/create',
  ORG_INVITATION_LIST: '/invite/list/org',

  ORG_MEMBERS_LIST: '/member',
  ORG_MEMBER_UPDATE_TYPE: '/member/update/type',
  ORG_MEMBER_DELETE: '/member/delete',
  ORG_MEMBER_TYPE: '/member/type',

  USER: '/user',
  USER_CREATE: '/user/create',
  USER_WITHDRAW: '/user/withdraw',

  CLI_USER: '/cli/user',
  CLI_USER_UPDATE: '/cli/user/update',
}
export type GetPath = typeof GetPath[keyof typeof GetPath]

type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE'

export const fetchBaseApi = async (
  path: string,
  method: HttpMethod,
  body?: { [key: string]: any },
  headers?: { [key: string]: string }
): Promise<Response> => {
  const opt: RequestInit = {
    method: method,
    mode: 'cors',
    credentials: 'include',
    headers: headers ?? {
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: body && JSON.stringify(body),
  }
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}${path}`, opt)

  switch (res.status) {
    case 401:
      // refresh
      await fetch(`${process.env.NEXT_PUBLIC_FRONT_URL}/api/auth/refresh`, {
        method: 'GET',
        credentials: 'include',
      })

      // re fetch
      const res2 = await fetch(`${process.env.NEXT_PUBLIC_API_URL}${path}`, opt)
      return res2
    default:
      return res
  }
}

export async function fetcherGetFromApiUrl(path: string) {
  try {
    const res = await fetchBaseApi(path, 'GET')
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
      const resUser = await fetch(
        `${process.env.NEXT_PUBLIC_FRONT_URL}/api/auth/server/refresh`,
        {
          method: 'GET',
          credentials: 'include',
        }
      )
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
