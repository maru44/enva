package mysmtp

import "github.com/maru44/enva/service/api/pkg/domain"

type ISmtpHandler interface {
	Send(domain.SmtpInput) error
}
