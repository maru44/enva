package usecase

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type OrgInvitationInteractor struct {
	repo IOrgInvitationRepository
}

func NewOrgInvitaionInteractor(repo IOrgInvitationRepository) domain.IOrgInvitationInteractor {
	return &OrgInvitationInteractor{
		repo: repo,
	}
}

type IOrgInvitationRepository interface {
	ListFromOrg(context.Context, domain.OrgID) ([]domain.OrgInvitation, error)
	List(context.Context) ([]domain.OrgInvitation, error)
	Detail(context.Context, domain.OrgInvitationID) (*domain.OrgInvitation, error)
	Create(context.Context, domain.OrgInvitationInput, domain.UserID) error
}

func (in *OrgInvitationInteractor) ListFromOrg(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitation, error) {
	return in.repo.ListFromOrg(ctx, orgID)
}

func (in *OrgInvitationInteractor) List(ctx context.Context) ([]domain.OrgInvitation, error) {
	return in.repo.List(ctx)
}

func (in *OrgInvitationInteractor) Detail(ctx context.Context, orgID domain.OrgInvitationID) (*domain.OrgInvitation, error) {
	return in.repo.Detail(ctx, orgID)
}

func (in *OrgInvitationInteractor) Create(ctx context.Context, input domain.OrgInvitationInput, userID domain.UserID) error {
	return in.repo.Create(ctx, input, userID)
}
