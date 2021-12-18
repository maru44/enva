package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type CommandInteractor struct {
	repo ICommandRepository
}

func NewCommandInteractor(repo ICommandRepository) domain.ICommandInteractor {
	return &CommandInteractor{
		repo: repo,
	}
}

type ICommandRepository interface {
	Run(ctx context.Context, topts ...string) error
}

func (in *CommandInteractor) Run(ctx context.Context, opts ...string) error {
	return in.repo.Run(ctx, opts...)
}
