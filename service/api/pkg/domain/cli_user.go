package domain

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/maru44/perr"
)

type (
	CliUser struct {
		CognitoID string
		Email     string
		Username  string
	}

	CliUserCreateInput struct {
		CognitoID string
		Email     string
		Username  string
		Password  string
	}

	CliUserValidateInput struct {
		EmailOrUsername string `json:"email_or_username"`
		Password        string `json:"password"`
	}

	ICliUserInteractor interface {
		Create(context.Context) (string, error)
		Update(context.Context) (string, error)
		Validate(context.Context, *CliUserValidateInput) error
		Exists(context.Context) error
		Delete(context.Context) error

		GetUser(context.Context, *CliUserValidateInput) (*User, error)
	}
)

func (c *CliUser) ToUser() *User {
	return &User{
		ID:       c.CognitoID,
		Email:    c.Email,
		UserName: c.Username,
	}
}

func (c *CliUserCreateInput) Validate() error {
	if err := validation.ValidateStruct(c,
		validation.Field(&c.CognitoID, validation.Required),
		validation.Field(&c.Password, validation.Required, validation.RuneLength(31, 255)),
		validation.Field(&c.Email, validation.Required, is.Email, validation.RuneLength(1, 255)),
		// validation.Field(&c.Username, validation.Required, validation.RuneLength(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}

func (c *CliUserValidateInput) Validate() error {
	if err := validation.ValidateStruct(c,
		validation.Field(&c.Password, validation.Required, validation.RuneLength(31, 255)),
		validation.Field(&c.EmailOrUsername, validation.Required, validation.RuneLength(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}

func CreateAndValidateCliUserCraeteInput(user User, hashed string) (*CliUserCreateInput, error) {
	fmt.Println("user id", user.ID)
	input := &CliUserCreateInput{
		CognitoID: user.ID,
		Email:     user.Email,
		Password:  hashed,
		Username:  user.UserName,
	}

	if err := input.Validate(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return input, nil
}

const (
	CLI_HEADER_SEP = "=+=+=+="
)
