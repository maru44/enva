package commands

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	get struct{}
)

func init() {
	Commands["get"] = func() domain.ICommandInteractor {
		return &get{}
	}
}

func (c *get) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// without keys
	if len(opts) == 0 {
		body, err := fetchListValid(ctx)
		if err != nil {
			return err
		}

		for _, d := range body.Data {
			color.Green(fmt.Sprintf("%s = %s", d.Key, d.Value))
		}

		return nil
	}

	// with keys

	// auth input
	email, password, err := inputEmailPassword()
	if err != nil {
		return err
	}

	kvs := []domain.KvValid{}

	for _, o := range opts {
		body, err := fetchDetailValid(ctx, o, email, password)
		if err != nil {
			return err
		}

		kvs = append(kvs, body.Data)
	}

	if kvs != nil {
		if len(opts) == 1 {
			color.Green(kvs[0].Value.String())
		} else {
			for _, kv := range kvs {
				color.Green("%s: %s\n", kv.Key, kv.Value)
			}
			fmt.Print("\n")
		}
	}

	return nil
}

func (c *get) Explain() string {
	return `	Getting remote key-value sets and output them in command line.
	If count of args is more than or equal 1, get value of the keys (key) designated in args.
	ex1) enva get
	ex2) enva get [keys...]
`
}
