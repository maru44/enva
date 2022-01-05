package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
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
	ListAll(context.Context) ([]domain.Project, error)
	ListByUser(context.Context) ([]domain.Project, error)
	ListByOrg(context.Context, domain.OrgID) ([]domain.Project, error)
	SlugListByUser(context.Context) ([]string, error)
	GetBySlug(context.Context, string) (*domain.Project, error)
	GetBySlugAndOrgID(context.Context, string, domain.OrgID) (*domain.Project, error)
	GetByID(context.Context, domain.ProjectID) (*domain.Project, error)
	Create(context.Context, domain.ProjectInput) (*string, error)
	Delete(context.Context, domain.ProjectID) (int, error)
}

func (in *ProjectInteractor) ListAll(ctx context.Context) ([]domain.Project, error) {
	return in.repo.ListAll(ctx)
}

func (in *ProjectInteractor) ListByUser(ctx context.Context) ([]domain.Project, error) {
	return in.repo.ListByUser(ctx)
}

func (in *ProjectInteractor) ListByOrg(ctx context.Context, orgID domain.OrgID) ([]domain.Project, error) {
	return in.repo.ListByOrg(ctx, orgID)
}

func (in *ProjectInteractor) SlugListByUser(ctx context.Context) ([]string, error) {
	return in.repo.SlugListByUser(ctx)
}

func (in *ProjectInteractor) GetBySlug(ctx context.Context, slug string) (*domain.Project, error) {
	return in.repo.GetBySlug(ctx, slug)
}

func (in *ProjectInteractor) GetBySlugAndOrgID(ctx context.Context, slug string, orgID domain.OrgID) (*domain.Project, error) {
	return in.repo.GetBySlugAndOrgID(ctx, slug, orgID)
}

func (in *ProjectInteractor) GetByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	return in.repo.GetByID(ctx, id)
}

func (in *ProjectInteractor) Create(ctx context.Context, input domain.ProjectInput) (*string, error) {
	return in.repo.Create(ctx, input)
}

func (in *ProjectInteractor) Delete(ctx context.Context, projectID domain.ProjectID) (int, error) {
	return in.repo.Delete(ctx, projectID)
}
