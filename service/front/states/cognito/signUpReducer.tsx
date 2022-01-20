export type signUpState = {
  username?: string
  email?: string
  password?: string
  password2?: string
  canSubmit: boolean
}

type signUpAction =
  | {
      type: 'setUsername'
      value?: string
    }
  | {
      type: 'setEmail'
      value?: string
    }
  | {
      type: 'setPassword'
      value?: string
    }
  | {
      type: 'setPassword2'
      value?: string
    }

export const initialSignUpState: signUpState = {
  username: '',
  email: '',
  password: '',
  password2: '',
  canSubmit: false,
}

export const signUpReducer = (
  state: signUpState,
  action: signUpAction
): signUpState => {
  switch (action.type) {
    case 'setEmail': {
      const st = { ...state, email: action.value }
      return { ...st, canSubmit: canSubmit(st) }
    }
    case 'setUsername': {
      const st = { ...state, username: action.value }
      return { ...st, canSubmit: canSubmit(st) }
    }
    case 'setPassword': {
      const st = { ...state, password: action.value }
      return { ...st, canSubmit: canSubmit(st) }
    }
    case 'setPassword2': {
      const st = { ...state, password2: action.value }
      return { ...st, canSubmit: canSubmit(st) }
    }
    default:
      return state
  }
}

const canSubmit = ({
  email,
  username,
  password,
  password2,
}: signUpState): boolean => {
  if (!email || !username || !password || !password2) return false
  if (email.length < 6) return false
  if (username.length < 4) return false
  if (password.length < 8) return false
  if (password2.length < 8) return false
  return password === password2
}
