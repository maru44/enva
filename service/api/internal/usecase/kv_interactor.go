package usecase

import (
	"context"

	"github.com/maru44/ichigo/service/api/pkg/domain"
)

type KvInteractor struct {
	repo IKvRepository
}

func NewKvInteractor(kv IKvRepository) domain.IKvInteractor {
	return &KvInteractor{
		repo: kv,
	}
}

// interface of kv repository
type IKvRepository interface {
	ListValid(context.Context, domain.ProjectID) ([]domain.Kv, error)
	DetailValid(context.Context, domain.KvKey, domain.ProjectID) (*domain.Kv, error)
	Create(context.Context, domain.KvInput, domain.ProjectID) (*domain.KvID, error)
	Update(context.Context, domain.KvInput, domain.ProjectID) (*domain.KvID, error)
}

/***********************************************
    implementation of kv interactor methods
************************************************/

func (in KvInteractor) ListValid(ctx context.Context, projectID domain.ProjectID) ([]domain.Kv, error) {
	return in.repo.ListValid(ctx, projectID)
}

func (in KvInteractor) DetailValid(ctx context.Context, key domain.KvKey, projectID domain.ProjectID) (*domain.Kv, error) {
	return in.repo.DetailValid(ctx, key, projectID)
}

func (in KvInteractor) Create(ctx context.Context, input domain.KvInput, projectID domain.ProjectID) (*domain.KvID, error) {
	return in.repo.Create(ctx, input, projectID)
}

func (in KvInteractor) Update(ctx context.Context, input domain.KvInput, projectID domain.ProjectID) (*domain.KvID, error) {
	return in.repo.Update(ctx, input, projectID)
}
