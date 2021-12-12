import { Theme } from '@mui/material'
import { ThemeName } from 'src/theme'

declare module '@mui/styles/defaultTheme' {
  interface DefaultTheme extends Theme {
    name: ThemeName
    light: boolean
  }
}

declare module '@mui/material/styles/createTheme' {
  export interface ThemeOptions {
    name: ThemeName
    light: boolean
  }

  export interface Theme {
    name: ThemeName
    light: boolean
  }
}

declare module '@mui/material/styles/adaptV4Theme' {
  export interface DeprecatedThemeOptions {
    name: ThemeName
    light: boolean
  }
}

declare module '@mui/material/styles/createPalette' {
  export interface TypeBackground {
    default: string
    paper: string
    dark: string
  }
}
