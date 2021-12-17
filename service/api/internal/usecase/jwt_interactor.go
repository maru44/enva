package usecase

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type JwtInteractor struct {
	repo IJwtRepository
}

func NewJwtInteractor(jr IJwtRepository) domain.JwtIntectactor {
	return &JwtInteractor{
		repo: jr,
	}
}

type IJwtRepository interface {
	Evaluate(context.Context, string) (*jwt.Token, error)
	GetUserByJwt(context.Context, string) (*domain.User, error)
}

func (in *JwtInteractor) GetUserByJwt(ctx context.Context, idToken string) (*domain.User, error) {
	return in.repo.GetUserByJwt(ctx, idToken)
}

func (in *JwtInteractor) Evaluate(ctx context.Context, idToken string) (*jwt.Token, error) {
	return in.repo.Evaluate(ctx, idToken)
}
