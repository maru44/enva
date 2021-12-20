package commands

import (
	"context"
	"errors"
	"fmt"

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

	fmt.Printf("%s = %s is added!\n", key, value)
	return nil
}

func (c *add) Explain() string {
	return `
	Add remote and local key-value set.
	Two args is needed. First arg is key, second arg is value.
	ex) enva add [key] [value]
`
}
