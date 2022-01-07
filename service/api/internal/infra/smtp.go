package infra

import (
	"fmt"
	"net/smtp"

	"github.com/maru44/enva/service/api/internal/config"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type Smtp struct{}

func (s *Smtp) Send(input domain.SmtpInput) error {
	if err := input.Validate(); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	smtpAuth := smtp.PlainAuth(
		"",
		config.EMAIL_HOST_USER,
		config.EMAIL_HOST_PASS,
		config.EMAIL_HOST,
	)

	if err := smtp.SendMail(
		config.EMAIL_HOST+":"+config.EMAIL_PORT,
		smtpAuth,
		config.EMAIL_HOST_USER,
		[]string{input.To},
		[]byte(fmt.Sprintf(
			`To: %s
Subject: %s
From: %s
Content-Type: text/plain; charset="utf-8"

%s
`, input.Subject, input.Subject, config.EMAIL_HOST_USER, input.Message,
		)),
	); err != nil {
		return perr.Wrap(err, perr.InternalServerError)
	}

	return nil
}
