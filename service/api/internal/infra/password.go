package infra

import (
	"github.com/maru44/perr"
	"golang.org/x/crypto/bcrypt"
)

type Password struct{}

func (p *Password) Check(hashed, raw string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(raw),
	)
}

func (p *Password) Generate(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), 8)
	if err != nil {
		return "", perr.Wrap(err, perr.InternalServerError)
	}

	return string(hashed), nil
}
