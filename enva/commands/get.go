package commands

import (
	"context"
	"fmt"

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
			fmt.Println(fmt.Sprintf("%s = %s", d.Key, d.Value))
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
	errs := map[string]error{}

	for _, o := range opts {
		body, err := fetchDetailValid(ctx, o, email, password)
		if err != nil {
			errs[o] = err
			continue
		}

		kvs = append(kvs, body.Data)
	}

	if kvs != nil {
		if len(opts) == 1 {
			fmt.Println(kvs[0].Value)
		} else {
			for _, kv := range kvs {
				fmt.Printf("%s: %s\n", kv.Key, kv.Value)
			}
			fmt.Print("\n")
		}
	}

	if errs != nil {
		for k, v := range errs {
			fmt.Printf("ERR: %s\n\t%s\n", k, v)
		}
	}

	return nil
}

func (c *get) Explain() string {
	return `
	Get remote key-value sets and output in command line.
	If count of args is larger than 1, get the value of keys designated in args.
	ex1) enva get
	ex2) enva get [key1] [key2]
`
}
