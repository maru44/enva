package commands

import (
	"context"
	"fmt"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/interface/password"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type (
	createCommandHandler func(sql database.ISqlHandler, pass password.IPassword) domain.ICommandInteractor
)

var Commands = map[string]createCommandHandler{}

func Run(ctx context.Context, sql database.ISqlHandler, pass password.IPassword, taskName string, opts ...string) error {
	createCommands, ok := Commands[taskName]
	if !ok {
		return perr.New("No such command", perr.BadRequest)
	}

	c := createCommands(sql, pass)
	if err := c.Run(ctx, opts...); err != nil {
		return err
	}

	fmt.Println("Done")
	return nil
}
