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
	SlugListByUser(context.Context) ([]string, error)
	Detail(context.Context, string) (*domain.Project, error)
	Create(context.Context, domain.ProjectInput) (*string, error)
}

func (in ProjectInteractor) ListByUser(ctx context.Context) ([]domain.Project, error) {
	return in.repo.ListByUser(ctx)
}

func (in ProjectInteractor) ListByOrg(ctx context.Context, orgID domain.OrgID) ([]domain.Project, error) {
	return in.repo.ListByOrg(ctx, orgID)
}

func (in ProjectInteractor) SlugListByUser(ctx context.Context) ([]string, error) {
	return in.repo.SlugListByUser(ctx)
}

func (in ProjectInteractor) Detail(ctx context.Context, slug string) (*domain.Project, error) {
	return in.repo.Detail(ctx, slug)
}

func (in ProjectInteractor) Create(ctx context.Context, input domain.ProjectInput) (*string, error) {
	return in.repo.Create(ctx, input)
}
