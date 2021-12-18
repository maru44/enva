package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type CliUserInteractor struct {
	repo ICliUserRepository
}

func NewCliUserInteractor(cr ICliUserRepository) domain.ICliUserInteractor {
	return &CliUserInteractor{
		repo: cr,
	}
}

type ICliUserRepository interface {
	Create(context.Context) (string, error)
	Update(context.Context) (string, error)
	Validate(context.Context, *domain.CliUserValidateInput) error
	Exists(context.Context) error
	Delete(context.Context) error

	GetUser(context.Context, *domain.CliUserValidateInput) (*domain.User, error)
}

/***********************************************
    implementation of cliuser interactor methods
************************************************/

func (in *CliUserInteractor) Create(ctx context.Context) (string, error) {
	return in.repo.Create(ctx)
}

func (in *CliUserInteractor) Update(ctx context.Context) (string, error) {
	return in.repo.Update(ctx)
}

func (in *CliUserInteractor) Validate(ctx context.Context, input *domain.CliUserValidateInput) error {
	return in.repo.Validate(ctx, input)
}

func (in *CliUserInteractor) Exists(ctx context.Context) error {
	return in.repo.Exists(ctx)
}

func (in *CliUserInteractor) Delete(ctx context.Context) error {
	return in.repo.Delete(ctx)
}

func (in *CliUserInteractor) GetUser(ctx context.Context, input *domain.CliUserValidateInput) (*domain.User, error) {
	return in.repo.GetUser(ctx, input)
}
