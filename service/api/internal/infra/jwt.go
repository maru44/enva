package infra

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/maru44/enva/service/api/internal/interface/myjwt"
)

type (
	JwtParser struct {
		Parser *jwt.Parser
	}

	JwtToken struct {
		Token *jwt.Token
	}
)

/***************************
    JwtParser instance
***************************/

func NewJwtParser() myjwt.JwtParserAbstract {
	parser := &jwt.Parser{
		SkipClaimsValidation: false,
	}

	jwtParser := new(JwtParser)
	jwtParser.Parser = parser
	return jwtParser
}

/***************************
     JwtParser methods
***************************/

func (p *JwtParser) Parse(tokenString string, keyFunc jwt.Keyfunc) (myjwt.JwtTokenAbstract, error) {
	jwtToken := new(JwtToken)
	t, err := p.Parser.Parse(tokenString, keyFunc)
	if err != nil {
		return jwtToken, err
	}
	jwtToken.Token = t
	return jwtToken, err
}

/***************************
     JwtToken methods
***************************/

func (t *JwtToken) GetToken() *jwt.Token {
	return t.Token
}
