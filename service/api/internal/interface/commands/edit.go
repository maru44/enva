package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	edit struct{}
)

func init() {
	Commands["edit"] = func() domain.ICommandInteractor {
		return &edit{}
	}
}

func (c *edit) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if opts == nil {
		return errors.New("Need two args key and value")
	}

	if len(opts) > 2 {
		return errors.New("Too many arguments")
	} else if len(opts) < 2 {
		return errors.New("Too few arguments.\nThe 'edit' command need key and value")
	}

	key := opts[0]
	value := opts[1]
	_, err := fetchUpdateKv(ctx, key, value)
	if err != nil {
		return err
	}

	if err := fileReadAndUpdateKv(key, value); err != nil {
		return err
	}

	fmt.Printf("%s = %s is updated!\n", key, value)
	return nil
}

func (c *edit) Explain() string {
	return `
	Edit remote and local value.
	Two args is needed. First arg is key, second arg is value.
	ex) enva edit [key] [value]`
}
