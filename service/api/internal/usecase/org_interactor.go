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
	ListOwnerAdmin(context.Context) ([]domain.Org, error)
	DetailBySlug(context.Context, string) (*domain.Org, *domain.UserType, error)
	Create(context.Context, domain.OrgInput) (*string, error)

	/* invitations */
	Invite(context.Context, domain.OrgInvitationInput) error
	InvitationListFromOrg(context.Context, domain.OrgID) ([]domain.OrgInvitation, error)
	InvitationList(context.Context) ([]domain.OrgInvitation, error)
	InvitationDetail(context.Context, domain.OrgInvitationID) (*domain.OrgInvitation, error)
	InvitationPastList(context.Context, domain.OrgID) ([]domain.OrgInvitationID, error)
	InvitationUpdateStatus(context.Context, domain.OrgInvitationID, domain.OrgInvitationStatus) error
	InvitationDeny(context.Context, domain.OrgInvitationID) error

	/* member */
	MemberCreate(context.Context, domain.OrgMemberInput) error
	MemberList(context.Context, domain.OrgID) (map[domain.UserType][]domain.User, error)
	MemberGetCurrentUserType(context.Context, domain.OrgID) (*domain.UserType, error)
	MemberGetUserType(context.Context, domain.UserID, domain.OrgID) (*domain.UserType, error)
	MemberUpdateUserType(context.Context, domain.OrgMemberUpdateInput) error
	MemberDelete(context.Context, domain.UserID, domain.OrgID) error
}

func (in *OrgInteractor) List(ctx context.Context) ([]domain.Org, error) {
	return in.repo.List(ctx)
}

func (in *OrgInteractor) ListOwnerAdmin(ctx context.Context) ([]domain.Org, error) {
	return in.repo.ListOwnerAdmin(ctx)
}

func (in *OrgInteractor) DetailBySlug(ctx context.Context, slug string) (*domain.Org, *domain.UserType, error) {
	return in.repo.DetailBySlug(ctx, slug)
}

func (in *OrgInteractor) Create(ctx context.Context, input domain.OrgInput) (*string, error) {
	return in.repo.Create(ctx, input)
}

/* invitation */

func (in *OrgInteractor) InvitationListFromOrg(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitation, error) {
	return in.repo.InvitationListFromOrg(ctx, orgID)
}

func (in *OrgInteractor) InvitationList(ctx context.Context) ([]domain.OrgInvitation, error) {
	return in.repo.InvitationList(ctx)
}

func (in *OrgInteractor) InvitationDetail(ctx context.Context, invID domain.OrgInvitationID) (*domain.OrgInvitation, error) {
	return in.repo.InvitationDetail(ctx, invID)
}

func (in *OrgInteractor) Invite(ctx context.Context, input domain.OrgInvitationInput) error {
	return in.repo.Invite(ctx, input)
}

func (in *OrgInteractor) InvitationPastList(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitationID, error) {
	return in.repo.InvitationPastList(ctx, orgID)
}

func (in *OrgInteractor) InvitationUpdateStatus(ctx context.Context, invID domain.OrgInvitationID, status domain.OrgInvitationStatus) error {
	return in.repo.InvitationUpdateStatus(ctx, invID, status)
}

func (in *OrgInteractor) InvitationDeny(ctx context.Context, invID domain.OrgInvitationID) error {
	return in.repo.InvitationDeny(ctx, invID)
}

/* member */

func (in *OrgInteractor) MemberCreate(ctx context.Context, input domain.OrgMemberInput) error {
	return in.repo.MemberCreate(ctx, input)
}

func (in *OrgInteractor) MemberList(ctx context.Context, orgID domain.OrgID) (map[domain.UserType][]domain.User, error) {
	return in.repo.MemberList(ctx, orgID)
}

func (in *OrgInteractor) MemberGetCurrentUserType(ctx context.Context, orgID domain.OrgID) (*domain.UserType, error) {
	return in.repo.MemberGetCurrentUserType(ctx, orgID)
}

func (in *OrgInteractor) MemberGetUserType(ctx context.Context, userID domain.UserID, orgID domain.OrgID) (*domain.UserType, error) {
	return in.repo.MemberGetUserType(ctx, userID, orgID)
}

func (in *OrgInteractor) MemberUpdateUserType(ctx context.Context, input domain.OrgMemberUpdateInput) error {
	return in.repo.MemberUpdateUserType(ctx, input)
}

func (in *OrgInteractor) MemberDelete(ctx context.Context, userID domain.UserID, orgID domain.OrgID) error {
	return in.repo.MemberDelete(ctx, userID, orgID)
}
