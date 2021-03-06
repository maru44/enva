package commands

import (
	"context"
	"errors"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	createCommandHandler func() domain.ICommandInteractor
)

var (
	Commands = map[string]createCommandHandler{}

	apiUrl = ""
)

// if added commands add the name to AllCommands
func Run(ctx context.Context, taskName string, opts ...string) {
	createCommands, ok := Commands[taskName]
	if !ok {
		err := errors.New("No such command")
		printErr(err)
		color.Yellow("\nIf you want to know commands, execute `enva help` command.")
		return
	}

	c := createCommands()
	if err := c.Run(ctx, opts...); err != nil {
		printErr(err)
	}
}

func printErr(err error) {
	color.Red("ERR: %s\n", err.Error())
}
