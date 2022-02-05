export type CurrentUser = {
  id: string
  username: string
  email: string
  is_valid: boolean
  image_url?: string
  is_email_verified: boolean
  has_cli_password: boolean
}

export const UserType = {
  OWNER: 'owner',
  ADMIN: 'admin',
  USER: 'user',
  GUEST: 'guest',
}
export type UserType = typeof UserType[keyof typeof UserType]

export const AdminUserTypes = [UserType.ADMIN, UserType.OWNER]

export const UserUserTypes = [UserType.ADMIN, UserType.OWNER, UserType.USER]

export const UserTypesAll: UserType[] = [
  UserType.OWNER,
  UserType.ADMIN,
  UserType.USER,
  UserType.GUEST,
]
