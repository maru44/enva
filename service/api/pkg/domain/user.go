package domain

import (
	"context"
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/maru44/perr"
)

type (
	UserID   string
	UserType string

	User struct {
		ID              UserID    `json:"id"`
		Username        string    `json:"username"`
		Email           string    `json:"email"`
		ImageURL        *string   `json:"image_url"`
		CliPassword     *string   `json:"-"`
		IsValid         bool      `json:"is_valid"`
		IsEmailVerified bool      `json:"is_email_verified"`
		CreatedAt       time.Time `json:"-"`
		UpdatedAt       time.Time `json:"-"`

		HasCliPassword bool `json:"has_cli_password"`
		// SshPubKeys []string `json:"ssh_pub_keys"`

		Subscription *Subscription `json:"subscription"`
	}

	UserInput struct {
		ID              string `json:"id"`
		Email           string `json:"email"`
		Username        string `json:"username"`
		IsEmailVerified bool   `json:"is_email_verified"`
	}

	UserCliPasswordInput struct {
		ID          string `json:"id"`
		CliPassword string `json:"password"`
	}

	UserCliValidationInput struct {
		EmailOrUsername string `json:"email_or_username"`
		CliPassword     string `json:"password"`
	}

	IUserInteractor interface {
		GetByID(context.Context, UserID) (*User, error)
		GetByEmail(context.Context, string) (*User, error)
		Create(context.Context) (*string, error)

		UpdateCliPassword(context.Context) (*string, error)
		GetUserCli(context.Context, *UserCliValidationInput) (*User, error)
	}
)

func (u *UserID) String() string {
	return string(*u)
}

func (u *UserInput) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.ID, validation.Required, is.UUID),
		validation.Field(&u.Email, validation.Required, is.Email, validation.RuneLength(1, 255)),
		validation.Field(&u.Username, validation.Required, validation.RuneLength(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}

func (u *UserCliPasswordInput) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.ID, validation.Required, is.UUID),
		validation.Field(&u.CliPassword, validation.Required, validation.RuneLength(31, 511)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}

func (u *UserCliValidationInput) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.CliPassword, validation.Required, validation.RuneLength(31, 255)),
		validation.Field(&u.EmailOrUsername, validation.Required, validation.RuneLength(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}

func UserFromCtx(ctx context.Context) (*User, error) {
	user, ok := ctx.Value(CtxUserKey).(User)
	if !ok {
		return nil, errors.New("No user in context")
	}
	return &user, nil
}

const (
	CLI_HEADER_SEP = "=+=+=+="
)
