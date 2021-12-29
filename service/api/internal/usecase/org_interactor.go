package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type OrgInteractor struct {
	repo IOrgRepository
}

func NewOrgInteractor(repo IOrgRepository) domain.IOrgInteractor {
	return &OrgInteractor{
		repo: repo,
	}
}

type IOrgRepository interface {
	List(context.Context) ([]domain.Org, error)
	Detail(context.Context, domain.OrgID) (*domain.Org, error)
	Create(context.Context, domain.OrgInput) (*string, error)
}

func (in *OrgInteractor) List(ctx context.Context) ([]domain.Org, error) {
	return in.repo.List(ctx)
}

func (in *OrgInteractor) Detail(ctx context.Context, orgID domain.OrgID) (*domain.Org, error) {
	return in.repo.Detail(ctx, orgID)
}

func (in *OrgInteractor) Create(ctx context.Context, input domain.OrgInput) (*string, error) {
	return in.repo.Create(ctx, input)
}
