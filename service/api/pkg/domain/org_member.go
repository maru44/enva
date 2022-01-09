package domain

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/maru44/perr"
)

type (
	OrgMemberInput struct {
		OrgID           OrgID           `json:"org_id"`
		UserID          UserID          `json:"user_id"`
		UserType        UserType        `json:"user_type"`
		OrgInvitationID OrgInvitationID `json:"org_invitation_id"`
	}

	OrgMemberUpdateInput struct {
		OrgID    OrgID     `json:"org_id"`
		UserID   UserID    `json:"user_id"`
		UserType *UserType `json:"user_type"`
	}
)

const (
	UserTypeOwner = UserType("owner")
	UserTypeAdmin = UserType("admin")
	UserTypeUser  = UserType("user")
	UserTypeGuest = UserType("guest")
)

func (u *UserType) IsAdmin() error {
	if *u == UserTypeOwner || *u == UserTypeAdmin {
		return nil
	}
	return perr.New("user is not admin or owner", perr.Forbidden, "you are not admin or owner of this org")
}

func (u *UserType) IsUser() error {
	if *u != UserTypeGuest {
		return nil
	}
	return perr.New("user is guest", perr.Forbidden, "you are guest of this org")
}

func (o *OrgInvitation) ToMemberInput() *OrgMemberInput {
	return &OrgMemberInput{
		OrgID:           o.Org.ID,
		UserID:          o.User.ID,
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

func (r *OrgMemberUpdateInput) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.OrgID, validation.Required, is.UUID),
		validation.Field(&r.UserID, validation.Required, is.UUID),
		validation.Field(&r.UserType, validation.Required, validation.In(UserTypeAdmin, UserTypeGuest, UserTypeOwner, UserTypeUser)),
	)
}
