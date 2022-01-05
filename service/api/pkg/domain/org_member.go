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

	IOrgMemberInteractor interface {
		Create(context.Context, OrgMemberInput) error
		List(context.Context, OrgID) (map[UserType][]User, error)
		// Update() // mainly userType
		// Delete()
		GetCurrentUserType(context.Context, OrgID) (*UserType, error)
	}
)

const (
	UserTypeOwner = UserType("owner")
	UserTypeAdmin = UserType("admin")
	UserTypeUser  = UserType("user")
	UserTypeGuest = UserType("guest")
)

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
