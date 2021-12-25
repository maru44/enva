export const slugify = (str: string): string => {
  return str
    .trim()
    .replace(/[^\w\s-]/g, '')
    .replace(/[\s]+/g, '-')
    .replace(/^-+|-+$/g, '')
}

export const isSlug = (str: string): boolean => {
  const regexExp = /^[a-z0-9-_]+$/gi

  return regexExp.test(str)
}
