import { ApiUrl } from '../config/env'
import { PostInput } from '../types/post'

export const fetchCreatePost = async (input: PostInput): Promise<Response> => {
  return await fetch(`${ApiUrl}/post/create/`, {
    method: 'POST',
    mode: 'cors',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    body: JSON.stringify(input),
  })
}
