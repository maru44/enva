package domain

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/maru44/perr"
)

type (
	OrgID string

	Org struct {
		ID          OrgID     `json:"id"`
		Slug        string    `json:"slug"`
		Name        string    `json:"name"`
		Description *string   `json:"description"`
		IsValid     bool      `json:"is_valid"`
		CreatedBy   User      `json:"created_by"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`

		// fk

		Admins []User `json:"admin"`
		Users  []User `json:"users"`
		Guests []User `json:"guests"`
	}

	OrgInput struct {
		Slug        string  `json:"slug"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}

	OrgMemberInput struct {
		OrgID           OrgID           `json:"org_id"`
		UserID          UserID          `json:"user_id"`
		UserType        UserType        `json:"user_type"`
		OrgInvitationID OrgInvitationID `json:"org_invitation_id"`
	}

	IOrgInteractor interface {
		List(context.Context) ([]Org, error)
		Detail(context.Context, string) (*Org, error)
		Create(context.Context, *OrgInput) (*OrgID, error)
		// Update
		// Delete
	}

	IOrgMemberInteractor interface {
		Create(context.Context, OrgMemberInput) error
		// Update() // mainly userType
		// Delete()
	}
)

func (o *OrgInput) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(o.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(o.Name, validation.Required, validation.Length(1, 64)),
	)
}

func (o *OrgInvitation) ToMemberInput() *OrgMemberInput {
	return &OrgMemberInput{
		OrgID:           o.OrgID,
		UserID:          o.UserID,
		UserType:        o.UserType,
		OrgInvitationID: o.ID,
	}
}

func (o *OrgMemberInput) Validate(ctx context.Context) error {
	user, err := UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}
	if user.ID != o.UserID {
		return perr.New("user not match: invited user and current user", perr.Forbidden)
	}

	return validation.ValidateStruct(o,
		validation.Field(&o.OrgID, validation.Required, is.UUID),
		validation.Field(&o.UserID, validation.Required, is.UUID),
		validation.Field(&o.OrgInvitationID, validation.Required, is.UUID),
		validation.Field(&o.UserType, validation.Required, validation.In(UserTypeOwner, UserTypeAdmin, UserTypeUser, UserTypeGuest)),
	)
}

func (o *Org) IsMember(u *User) bool {
	if u == nil {
		return false
	}
	for _, user := range o.Users {
		if user.ID == u.ID {
			return true
		}
	}

	return o.IsAdmin(u)
}

func (o *Org) IsAdmin(u *User) bool {
	if u == nil {
		return false
	}
	for _, user := range o.Admins {
		if user.ID == u.ID {
			return true
		}
	}

	return false
}
