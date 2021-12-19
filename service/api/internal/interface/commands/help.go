package commands

import (
	"context"
	"fmt"

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
	for name, in := range Commands {
		fmt.Print(name)
		f := in()
		fmt.Println(f.Explain())
	}
	return nil
}

func (c *help) Explain() string {
	return `
	Output explain how to use this cli.`
}
