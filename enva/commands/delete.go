package commands

import (
	"context"
	"errors"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	deleteS struct{}
)

func init() {
	Commands["delete"] = func() domain.ICommandInteractor {
		return &deleteS{}
	}
}

func (c *deleteS) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if len(opts) > 1 {
		return errors.New("Too many arguments")
	} else if len(opts) < 1 {
		return errors.New("Too few arguments.\nThe 'delete' command need key")
	}

	key := opts[0]
	_, err := fetchDeleteKv(ctx, key)
	if err != nil {
		return err
	}

	if err := fileReadAndDeleteKv(key); err != nil {
		return err
	}

	color.Green("%s is deleted!", key)
	return nil
}

func (c *deleteS) Explain() string {
	return `	Removing remote and local key-value set. An arg is needed.
	ex) enva delete [key]
`
}
