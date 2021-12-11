package usecase

import (
	"context"

	"github.com/maru44/ichigo/service/api/pkg/domain"
)

type ProjectInteractor struct {
	repo IProjectRepository
}

func NewProjectInteractor(repo IProjectRepository) domain.IProjectInteractor {
	return &ProjectInteractor{
		repo: repo,
	}
}

type IProjectRepository interface {
	ListByUser(context.Context) ([]domain.Project, error)
	ListByOrg(context.Context, domain.OrgID) ([]domain.Project, error)
	Detail(context.Context, domain.ProjectID) (*domain.Project, error)
	Create(context.Context, domain.ProjectInput) (*domain.ProjectID, error)
}

func (in ProjectInteractor) ListByUser(ctx context.Context) ([]domain.Project, error) {
	return in.repo.ListByUser(ctx)
}

func (in ProjectInteractor) ListByOrg(ctx context.Context, orgID domain.OrgID) ([]domain.Project, error) {
	return in.repo.ListByOrg(ctx, orgID)
}

func (in ProjectInteractor) Detail(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	return in.repo.Detail(ctx, id)
}

func (in ProjectInteractor) Create(ctx context.Context, input domain.ProjectInput) (*domain.ProjectID, error) {
	return in.repo.Create(ctx, input)
}
