package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	createCommandHandler func() domain.ICommandInteractor
)

var Commands = map[string]createCommandHandler{}

func Run(ctx context.Context, taskName string, opts ...string) error {
	createCommands, ok := Commands[taskName]
	if !ok {
		err := errors.New("No such command")
		fmt.Println(err)
		return err
	}

	c := createCommands()
	if err := c.Run(ctx, opts...); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
