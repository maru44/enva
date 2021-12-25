package domain

import (
	"context"
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
		CliPassword     *string   `json:"password"`
		IsValid         bool      `json:"is_valid"`
		IsEmailVerified bool      `json:"is_email_verified"`
		UserType        *UserType `json:"user_type"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updatd_at"`

		// fk

		MemberOf []ProjectID `json:"member_of"`
		AdminOf  []ProjectID `json:"admin_of"`
		GuestOf  []ProjectID `json:"guest_of"`
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
		Create(context.Context) (*string, error)

		UpdateCliPassword(context.Context) (*string, error)
		GetUserCli(context.Context, *UserCliValidationInput) (*User, error)
	}
)

const (
	UserTypeAdmin = UserType("admin")
	UserTypeUser  = UserType("user")
	UserTypeGuest = UserType("guest")
)

func (u *User) IsAdmin() bool {
	if u.UserType == nil {
		return false
	}
	return *u.UserType == UserTypeAdmin
}

func (u *UserID) String() string {
	return string(*u)
}

func (u *UserInput) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.ID, validation.Required, is.UUID),
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(1, 255)),
		validation.Field(&u.Username, validation.Required, validation.Length(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}

func (u *UserCliPasswordInput) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.ID, validation.Required, is.UUID),
		validation.Field(&u.CliPassword, validation.Required, validation.Length(31, 511)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}

func (u *UserCliValidationInput) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.CliPassword, validation.Required, validation.Length(31, 255)),
		validation.Field(&u.EmailOrUsername, validation.Required, validation.Length(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}

const (
	CLI_HEADER_SEP = "=+=+=+="
)
