package commands

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

type (
	help struct{}
)

func (c *help) Run(ctx context.Context, opts ...string) error {
	fmt.Print("\n==============================\n== How to use enva commands ==\n==============================\n\n")
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
