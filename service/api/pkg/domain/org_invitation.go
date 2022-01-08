package domain

import (
	"context"
	"fmt"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/maru44/enva/service/api/internal/config"
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
		OrgID    OrgID  `json:"org_id"`
		OrgName  string `json:"org_name"`
		User     *User
		Eamil    string   `json:"email"`
		UserType UserType `json:"user_type"`
	}

	IOrgInvitationInteractor interface {
		Create(context.Context, OrgInvitationInput) error
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
		Deny(context.Context, OrgInvitationID) error
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
		validation.Field(&o.OrgID, validation.Required, is.UUID),
		validation.Field(&o.OrgName, validation.Required),
		validation.Field(&o.Eamil, validation.Required, is.Email),
		validation.Field(&o.UserType, validation.Required, validation.In(UserTypeOwner, UserTypeAdmin, UserTypeUser, UserTypeGuest)),
	)
}

func (o *OrgInvitationInput) CreateMail(invID OrgInvitationID, inviter User) SmtpInput {
	subject := "Invitation to " + o.OrgName

	var message string
	if o.User != nil {
		message = fmt.Sprintf(`Dear %s

You are invited to an org name is %s as '%s' type user from %s. 
Click following link and accept or deny this invitation.

%s

Thank you.
`,
			o.User.Username, o.OrgName, o.UserType,
			inviter.Username, config.FRONT_URL+"/org/invite/"+string(invID),
		)
	} else {
		message = fmt.Sprintf(`Dear %s

You are invited to an org name is %s as '%s' type user from %s.
Maybe you have not created an account.

If you want to join this org, you have to create an account.
Click following link and create an account, then accept this invitation.

%s

Thank you.
`,
			o.Eamil, o.OrgName, o.UserType,
			inviter.Username,
			config.FRONT_URL+"/org/invite/"+string(invID),
		)
	}

	return SmtpInput{
		Subject: subject,
		Message: message,
		To:      o.Eamil,
	}
}
