package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type UserInteractor struct {
	repo IUserRepository
}

func NewUserInteractor(repo IUserRepository) domain.IUserInteractor {
	return &UserInteractor{
		repo: repo,
	}
}

type IUserRepository interface {
	GetByID(context.Context, domain.UserID) (*domain.User, error)
	GetByEmail(context.Context, string) (*domain.User, error)
	Create(context.Context) (*string, error)

	UpdateCliPassword(context.Context) (*string, error)
	// used in cli
	GetUserCli(context.Context, *domain.UserCliValidationInput) (*domain.User, error)
}

func (in *UserInteractor) GetByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	return in.repo.GetByID(ctx, id)
}

func (in *UserInteractor) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	return in.repo.GetByEmail(ctx, email)
}

func (in *UserInteractor) Create(ctx context.Context) (*string, error) {
	return in.repo.Create(ctx)
}

func (in *UserInteractor) UpdateCliPassword(ctx context.Context) (*string, error) {
	return in.repo.UpdateCliPassword(ctx)
}

func (in *UserInteractor) GetUserCli(ctx context.Context, input *domain.UserCliValidationInput) (*domain.User, error) {
	return in.repo.GetUserCli(ctx, input)
}
