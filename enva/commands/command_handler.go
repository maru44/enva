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

func init() {
	Commands["pull"] = func() domain.ICommandInteractor {
		return &pull{}
	}

	Commands["get"] = func() domain.ICommandInteractor {
		return &get{}
	}

	Commands["add"] = func() domain.ICommandInteractor {
		return &add{}
	}

	Commands["edit"] = func() domain.ICommandInteractor {
		return &edit{}
	}

	Commands["delete"] = func() domain.ICommandInteractor {
		return &delete{}
	}

	Commands["diff"] = func() domain.ICommandInteractor {
		return &diff{}
	}

	Commands["init"] = func() domain.ICommandInteractor {
		return &initialize{}
	}

	Commands["set"] = func() domain.ICommandInteractor {
		return &set{}
	}

	Commands["help"] = func() domain.ICommandInteractor {
		return &help{}
	}

	Commands["version"] = func() domain.ICommandInteractor {
		return &version{}
	}
}

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
