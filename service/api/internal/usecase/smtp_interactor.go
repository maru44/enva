package usecase

import "github.com/maru44/enva/service/api/pkg/domain"

type SmtpInteractor struct {
	repo ISmtpReposiotry
}

func NewSmtpInteractor(si ISmtpReposiotry) domain.ISmtpInteractor {
	return &SmtpInteractor{
		repo: si,
	}
}

type ISmtpReposiotry interface {
	Send(domain.SmtpInput) error
}

func (in *SmtpInteractor) Send(input domain.SmtpInput) error {
	return in.repo.Send(input)
}
