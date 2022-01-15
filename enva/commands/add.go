package commands

import (
	"context"
	"errors"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	add struct{}
)

func init() {
	Commands["add"] = func() domain.ICommandInteractor {
		return &add{}
	}
}

func (c *add) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if opts == nil {
		return errors.New("Need two arguments (key and value)")
	}

	if len(opts) > 2 {
		return errors.New("Too many arguments")
	} else if len(opts) < 2 {
		return errors.New("Too few arguments")
	}

	key := opts[0]
	value := opts[1]
	_, err := fetchCreateKv(ctx, key, value)
	if err != nil {
		return err
	}

	if err := fileReadAndUpdateKv(key, value); err != nil {
		return err
	}

	color.Green("%s = %s is added!", key, value)
	return nil
}

func (c *add) Explain() string {
	return `
	Adding remote and local key-value set.
	Two args are required. First one is key, another is value.
	ex) enva add [key] [value]
`
}
