package commands

import (
	"context"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	pull struct{}
)

func init() {
	Commands["pull"] = func() domain.ICommandInteractor {
		return &pull{}
	}
}

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

	return nil
}
