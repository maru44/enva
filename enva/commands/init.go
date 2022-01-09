package commands

import (
	"context"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	initialize struct{}
)

func init() {
	Commands["init"] = func() domain.ICommandInteractor {
		return &initialize{}
	}
}

func (c *initialize) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	kvs, err := fileReadAndCreateKvs()
	if err != nil {
		return err
	}

	inputs := make([]domain.KvInput, len(kvs))
	for i, kv := range kvs {
		input := kv.ToInput()
		if err := input.Validate(); err != nil {
			return err
		}
		inputs[i] = *input
	}

	if _, err := fetchBulkInsertKvs(ctx, inputs); err != nil {
		return err
	}

	color.Green("init project is succeded")
	return nil
}

func (c *initialize) Explain() string {
	return `
	Set key-value sets of remote based on local env file.
	This command is so powerful that you can't execute if any remote key-value is set in the project.
`
}
