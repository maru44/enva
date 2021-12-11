// export const IsDevelopment = process.env.NEXT_PUBLIC_ENV === 'dev'
export const IsDevelopment = process.env.NODE_ENV === 'development'
export const IsTest = process.env.NODE_ENV === 'test'
export const IsProduction = process.env.NODE_ENV === 'production'

export const ApiUrl = IsProduction ? '' : 'http://localhost:8080'
export const ThisUrl = IsProduction ? '' : 'http://localhost:3000'