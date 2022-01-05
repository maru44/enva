package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type CliKvInteractor struct {
	repo ICliKvRepository
}

func NewCliKvInteractor(repo ICliKvRepository) domain.ICliKvInteractor {
	return &CliKvInteractor{
		repo: repo,
	}
}

type ICliKvRepository interface {
	BulkInsert(context.Context, domain.ProjectID, []domain.KvInput) error
}

func (in *CliKvInteractor) BulkInsert(ctx context.Context, projectID domain.ProjectID, inputs []domain.KvInput) error {
	return in.repo.BulkInsert(ctx, projectID, inputs)
}
