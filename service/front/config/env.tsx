export const IsDevelopment = process.env.NODE_ENV === 'development'
export const IsTest = process.env.NODE_ENV === 'test'
export const IsProduction = process.env.NODE_ENV === 'production'

export const ApiUrl = IsProduction
  ? 'https://api.envassador.com'
  : 'http://localhost:8080'
export const ThisUrl = IsProduction
  ? 'https://envassador.com'
  : 'http://localhost:3000'
