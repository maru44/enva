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
		Email           string  `json:"email"`
		EmailVerified   bool    `json:"email_verified"`
		CognitoUserName string  `json:"cognito:username"`
		Sub             string  `json:"sub"`
		Picture         *string `json:"picture,omitempty"`
	}
)

const (
	JwtCookieKeyIdToken      = "id_token"
	JwtCookieKeyAccessToken  = "access_token"
	JwtCookieKeyRefreshToken = "refresh_token"
)

func (u *UserFromClaim) ToUser() *User {
	return &User{
		ID:              UserID(u.Sub),
		Email:           u.Email,
		Username:        u.CognitoUserName,
		IsEmailVerified: u.EmailVerified,
		ImageURL:        u.Picture,
	}
}
