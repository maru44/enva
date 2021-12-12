import { adaptV4Theme, DeprecatedThemeOptions, Theme } from '@mui/material'
import { createTheme } from '@mui/system'

export enum ThemeName {
  Light = 'LIGHT',
}

// export const createMyTheme = (name: ThemeName): Theme => {
//   const themeOptions = themesOptions.find((theme) => theme.name === name)
//   if (!themeOptions) {
//     // @TODO throw exceptions
//   }
//   const ret =  createTheme(adaptV4Theme(themeOptions))
// }

// const themesOptions: DeprecatedThemeOptions[] = [
//   {
//     name: ThemeName.Light,
//     light: true,
//     direction: 'ltr',
//     // typography,
//     // overrides,
//     // palette,
//     // shadows: softShadows,
//   },
// ]
