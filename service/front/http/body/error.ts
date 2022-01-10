export type errorResponseBody = {
  error: string
  status: number
}

export const internalServerErrorInFront = (): errorResponseBody => {
  return {
    error: 'Internal Server Error',
    status: 500,
  }
}
