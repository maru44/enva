package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	SmtpInput struct {
		Subject string
		Message string
		// should be email
		To string
	}

	ISmtpInteractor interface {
		Send(SmtpInput) error
	}
)

func (s *SmtpInput) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Subject, validation.Required),
		validation.Field(&s.Message, validation.Required),
		validation.Field(&s.To, validation.Required, is.Email),
	)
}
