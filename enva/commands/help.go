package commands

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	help struct{}
)

var AllCommands = []string{
	"pull", "get", "add", "edit",
	"delete", "diff", "init", "set",
	// "help", "version",
}

func init() {
	Commands["help"] = func() domain.ICommandInteractor {
		return &help{}
	}
}

func (c *help) Run(ctx context.Context, opts ...string) error {
	fmt.Print("\n==============================\n== How to use enva commands ==\n==============================\n\n")
	for _, name := range AllCommands {
		color.Green(name)
		cmd := Commands[name]
		fmt.Println(cmd().Explain())
	}
	fmt.Print("\n\n")
	return nil
}

func (c *help) Explain() string {
	return `
	Help! I need somebody.
	Help! Not just anybody.
`
}
