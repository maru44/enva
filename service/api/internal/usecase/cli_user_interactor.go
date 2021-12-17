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
	Create(context.Context, *domain.CliUserCreateInput) (string, error)
	Update(context.Context, *domain.CliUserCreateInput) (string, error)
	Validate(context.Context, *domain.CliUserValidateInput) error
	Delete(context.Context, string) error // 2nd arg is email or username
}

/***********************************************
    implementation of cliuser interactor methods
************************************************/

func (in *CliUserInteractor) Create(ctx context.Context, input *domain.CliUserCreateInput) (string, error) {
	return in.repo.Create(ctx, input)
}

func (in *CliUserInteractor) Update(ctx context.Context, input *domain.CliUserCreateInput) (string, error) {
	return in.repo.Update(ctx, input)
}

func (in *CliUserInteractor) Validate(ctx context.Context, input *domain.CliUserValidateInput) error {
	return in.repo.Validate(ctx, input)
}

func (in *CliUserInteractor) Delete(ctx context.Context, emailOrUsername string) error {
	return in.repo.Delete(ctx, emailOrUsername)
}
