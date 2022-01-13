package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	OrgID string

	Org struct {
		ID          OrgID      `json:"id"`
		Slug        string     `json:"slug"`
		Name        string     `json:"name"`
		Description *string    `json:"description"`
		IsValid     bool       `json:"is_valid"`
		CreatedBy   User       `json:"created_by"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`

		UserCount int `json:"user_count"`

		Subscription *Subscription `json:"subscription"`
	}

	OrgInput struct {
		Slug        string  `json:"slug"`
		Name        string  `json:"name"`
		Description *string `json:"description,omitempty"`
	}

	IOrgInteractor interface {
		List(context.Context) ([]Org, error)
		ListOwnerAdmin(context.Context) ([]Org, error)
		DetailBySlug(context.Context, string) (*Org, *UserType, error)
		Create(context.Context, OrgInput) (*string, error)
		OrgValidCount(context.Context, UserID) (*int, *Subscription, error)

		/* invitations */
		Invite(context.Context, OrgInvitationInput) error
		InvitationListFromOrg(context.Context, OrgID) ([]OrgInvitation, error)
		InvitationList(context.Context) ([]OrgInvitation, error)
		InvitationDetail(context.Context, OrgInvitationID) (*OrgInvitation, error)
		InvitationPastList(context.Context, OrgID) ([]OrgInvitationID, error)
		InvitationUpdateStatus(context.Context, OrgInvitationID, OrgInvitationStatus) error
		InvitationDeny(context.Context, OrgInvitationID) error

		/* member */
		MemberCreate(context.Context, OrgMemberInput) error
		MemberList(context.Context, OrgID) (map[UserType][]User, error)
		MemberGetCurrentUserType(context.Context, OrgID) (*UserType, error)
		MemberGetUserType(context.Context, UserID, OrgID) (*UserType, error)
		MemberUpdateUserType(context.Context, OrgMemberUpdateInput) error
		MemberDelete(context.Context, UserID, OrgID) error
		MemberValidCount(context.Context, OrgID) (*int, *Subscription, error)
	}
)

func (o *OrgInput) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.Slug, validation.Required, validation.Length(1, 64)),
		validation.Field(&o.Name, validation.Required, validation.Length(1, 64)),
	)
}

var (
	ErrOrgIsNotValid = errors.New("Org is not valid")
)

func (o *Org) Valid() error {
	if !o.IsValid || o.DeletedAt != nil {
		return ErrOrgIsNotValid
	}
	return nil
}

func CanCreateOrg(s *Subscription, orgCount int) error {
	errStr := "Orgs count reaches maximum (%d)"

	if s == nil && orgCount > 0 {
		return fmt.Errorf(errStr, 1)
	}
	// if s.StripeProductID == a {}
	// if s.StripeProductID == b {}

	return nil
}

func CanCreateOrgMember(s *Subscription, memberCount int) error {
	errStr := "Members count reaches maximum (%d)"

	if s == nil && memberCount > 4 {
		return fmt.Errorf(errStr, 5)
	}
	// if s.StripeProductID == a {}
	// if s.StripeProductID == b {}

	return nil
}
