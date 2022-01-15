package commands

import (
	"context"
	"errors"
	"os"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	createCommandHandler func() domain.ICommandInteractor
)

var (
	Commands = map[string]createCommandHandler{}

	ApiUrl = os.Getenv("CLI_API_URL")
)

// add to command is written in z_initialize.go
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
