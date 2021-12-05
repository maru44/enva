package myjwt

import "github.com/golang-jwt/jwt/v4"

type JwtParserAbstract interface {
	Parse(tokenString string, keyFunc jwt.Keyfunc) (JwtTokenAbstract, error)
}

type JwtTokenAbstract interface {
	GetToken() *jwt.Token
}
