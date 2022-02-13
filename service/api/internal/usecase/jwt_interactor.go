package usecase

import (
	"context"

	"github.com/lestrrat-go/jwx/jwk"
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
	GetUserByJwt(context.Context, string) (*domain.User, error)
	FetchJwk(context.Context, string) (jwk.Set, error)
}

func (in *JwtInteractor) GetUserByJwt(ctx context.Context, idToken string) (*domain.User, error) {
	return in.repo.GetUserByJwt(ctx, idToken)
}

func (in *JwtInteractor) FetchJwk(ctx context.Context, url string) (jwk.Set, error) {
	return in.repo.FetchJwk(ctx, url)
}
