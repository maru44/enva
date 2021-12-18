package commands

import (
	"context"
	"fmt"

	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

// @TODO sql and pass is not needed

type (
	createCommandHandler func() domain.ICommandInteractor
)

var Commands = map[string]createCommandHandler{}

func Run(ctx context.Context, taskName string, opts ...string) error {
	createCommands, ok := Commands[taskName]
	if !ok {
		return perr.New("No such command", perr.BadRequest)
	}

	c := createCommands()
	if err := c.Run(ctx, opts...); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
