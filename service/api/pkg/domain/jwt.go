package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type (
	JwtIntectactor interface {
		Evaluate(context.Context, string) (*jwt.Token, error)
		GetUserByJwt(context.Context, string) (*User, error)
	}

	UserFromClaim struct {
		Email           string `json:"email"`
		EmailVerified   bool   `json:"email_verified"`
		CognitoUserName string `json:"cognito:username"`
	}
)

const (
	JwtCookieKeyIdToken      = "id_token"
	JwtCookieKeyAccessToken  = "access_token"
	JwtCookieKeyRefreshToken = "refresh_token"
)

func (uc *UserFromClaim) ToUser() *User {
	return &User{
		ID:              uc.CognitoUserName,
		Email:           uc.Email,
		IsEmailVerified: uc.EmailVerified,
	}
}
