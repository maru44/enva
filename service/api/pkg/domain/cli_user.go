package domain

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/maru44/perr"
)

type (
	CliUserCreateInput struct {
		Email    string
		Username string
		Password string
	}

	CliUserValidateInput struct {
		EmailOrUsername string
		Password        string
	}

	ICliUserInteractor interface {
		Create(context.Context, *CliUserCreateInput) (string, error)
		Update(context.Context, *CliUserCreateInput) (string, error)
		Validate(context.Context, *CliUserValidateInput) error
		Delete(context.Context, string) error // 2nd arg is email or username
	}
)

func (c *CliUserCreateInput) Validate() error {
	if err := validation.ValidateStruct(&c,
		validation.Field(&c.Password, validation.Required, validation.RuneLength(31, 255)),
		validation.Field(&c.Email, validation.Required, is.Email, validation.RuneLength(1, 255)),
		validation.Field(&c.Username, validation.Required, validation.RuneLength(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	return nil
}

func (c *CliUserValidateInput) Validate() error {
	if err := validation.ValidateStruct(&c,
		validation.Field(&c.Password, validation.Required, validation.RuneLength(31, 255)),
		validation.Field(&c.EmailOrUsername, validation.Required, validation.RuneLength(1, 255)),
	); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}
