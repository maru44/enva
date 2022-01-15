package commands

import (
	"context"

	"github.com/fatih/color"
)

type (
	pull struct{}
)

func (c *pull) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	body, err := fetchListValid(ctx)
	if err != nil {
		return err
	}

	if err := fileWriteFromResponse(*body); err != nil {
		return err
	}

	color.Green("\nSucceeded to pull remote env!\n\n")

	return nil
}

func (c *pull) Explain() string {
	return `
	Pulling remote key-value sets and overwrite your env file written in enva.json.
`
}
