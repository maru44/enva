package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	delete struct{}
)

func init() {
	Commands["delete"] = func() domain.ICommandInteractor {
		return &delete{}
	}
}

func (c *delete) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if opts == nil {
		return errors.New("Need one an arg")
	}

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

	fmt.Printf("%s is deleted!\n", key)
	return nil
}

func (c *delete) Explain() string {
	return `
	Remove remote and local key-value set.
	An arg is needed.
	ex) enva delete [key]
`
}
