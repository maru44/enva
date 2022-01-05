package domain

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	OrgInvitationID string

	OrgInvitation struct {
		ID       OrgInvitationID `json:"id"`
		UserType UserType        `json:"user_type"`
		IsValid  bool            `json:"is_valid"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`

		Org     Org  `json:"org"`
		User    User `json:"user"`
		Invitor User `json:"invitor"`
	}

	OrgInvitationInput struct {
		OrgID    OrgID    `json:"org_id"`
		UserID   UserID   `json:"user_id"`
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
	}
)

func (o *OrgInvitationInput) Validate() error {
	return validation.ValidateStruct(o,
		validation.Field(&o.OrgID, validation.Required, is.UUID),
		validation.Field(&o.UserID, validation.Required, is.UUID),
		validation.Field(&o.UserType, validation.Required, validation.In(UserTypeOwner, UserTypeAdmin, UserTypeUser, UserTypeGuest)),
	)
}

// func (o *OrgInvitation) ToMember()
