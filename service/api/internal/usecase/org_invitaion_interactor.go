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
	Create(context.Context, domain.OrgInvitationInput) error
	ListPastInvitations(context.Context, domain.OrgID) ([]domain.OrgInvitationID, error)
	UpdateStatus(context.Context, domain.OrgInvitationID, domain.OrgInvitationStatus) error
}

func (in *OrgInvitationInteractor) ListFromOrg(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitation, error) {
	return in.repo.ListFromOrg(ctx, orgID)
}

func (in *OrgInvitationInteractor) List(ctx context.Context) ([]domain.OrgInvitation, error) {
	return in.repo.List(ctx)
}

func (in *OrgInvitationInteractor) Detail(ctx context.Context, invID domain.OrgInvitationID) (*domain.OrgInvitation, error) {
	return in.repo.Detail(ctx, invID)
}

func (in *OrgInvitationInteractor) Create(ctx context.Context, input domain.OrgInvitationInput) error {
	return in.repo.Create(ctx, input)
}

func (in *OrgInvitationInteractor) ListPastInvitations(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitationID, error) {
	return in.repo.ListPastInvitations(ctx, orgID)
}

func (in *OrgInvitationInteractor) UpdateStatus(ctx context.Context, invID domain.OrgInvitationID, status domain.OrgInvitationStatus) error {
	return in.repo.UpdateStatus(ctx, invID, status)
}
