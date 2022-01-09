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

func init() {
	Commands["help"] = func() domain.ICommandInteractor {
		return &help{}
	}
}

func (c *help) Run(ctx context.Context, opts ...string) error {
	fmt.Print("\n\n\n\n\n")
	for name, in := range Commands {
		color.Green(name)
		f := in()
		fmt.Println(f.Explain())
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
