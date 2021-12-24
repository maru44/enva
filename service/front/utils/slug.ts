export const slugify = (str: string): string => {
  return str
    .toLowerCase()
    .trim()
    .replace(/[^\w\s-]/g, '')
    .replace(/[\s_-]+/g, '-')
    .replace(/^-+|-+$/g, '')
}

export const isSlug = (str: string): boolean => {
  const regexExp = /^[a-z0-9-_]+$/gi

  return regexExp.test(str)
}
