package domain

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	OrgInvitationID     string
	OrgInvitationStatus string

	OrgInvitation struct {
		ID       OrgInvitationID     `json:"id"`
		UserType UserType            `json:"user_type"`
		Status   OrgInvitationStatus `json:"status"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		Org     Org  `json:"org"`
		User    User `json:"user"`
		Invitor User `json:"invitor"`
	}

	OrgInvitationInput struct {
		OrgID    OrgID    `json:"org_id"`
		UserID   *UserID  `json:"user_id,omitempty"`
		Eamil    string   `json:"email"`
		UserType UserType `json:"user_type"`
	}

	IOrgInvitationInteractor interface {
		Create(context.Context, OrgInvitationInput, UserID) error
		// sent from
		ListFromOrg(context.Context, OrgID) ([]OrgInvitation, error)
		// sent by anyone
		List(context.Context) ([]OrgInvitation, error)
		// detail
		Detail(context.Context, OrgInvitationID) (*OrgInvitation, error)
		// past invitation to an email (from an org)
		ListPastInvitations(context.Context, OrgID) ([]OrgInvitationID, error)
		// Update status
		UpdateStatus(context.Context, OrgInvitationID, OrgInvitationStatus) error
	}
)

const (
	OrgInvitationStatusNew      = OrgInvitationStatus("new")
	OrgInvitationStatusAccepted = OrgInvitationStatus("accepted")
	OrgInvitationStatusDenied   = OrgInvitationStatus("denied")
	OrgInvitationStatusClosed   = OrgInvitationStatus("closed")
)

func (o *OrgInvitationInput) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.OrgID, is.UUID),
		validation.Field(&o.Eamil, validation.Required, is.Email),
		validation.Field(&o.UserID, validation.Required, is.UUID),
		validation.Field(&o.UserType, validation.Required, validation.In(UserTypeOwner, UserTypeAdmin, UserTypeUser, UserTypeGuest)),
	)
}

// func (o *OrgInvitation) ToMember()
