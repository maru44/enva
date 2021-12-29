package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type OrgMemberInteractor struct {
	repo IOrgMemberReposiotry
}

func NewOrgMemberInteractor(repo IOrgMemberReposiotry) domain.IOrgMemberInteractor {
	return &OrgMemberInteractor{
		repo: repo,
	}
}

type IOrgMemberReposiotry interface {
	Create(context.Context, domain.OrgMemberInput) error
}

func (in *OrgMemberInteractor) Create(ctx context.Context, input domain.OrgMemberInput) error {
	return in.repo.Create(ctx, input)
}
