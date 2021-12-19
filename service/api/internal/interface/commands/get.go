package commands

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	kvList struct{}
)

func init() {
	Commands["get"] = func() domain.ICommandInteractor {
		return &kvList{}
	}
}

func (c *kvList) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// with keys
	if opts != nil {
		// auth input
		input, err := inputEmailPassword()
		if err != nil {
			return err
		}
		inputJ, err := json.Marshal(input)
		if err != nil {
			return err
		}

		kvs := []domain.KvValid{}
		errs := map[string]error{}

		for _, o := range opts {
			body, err := fetchDetailValid(ctx, o, inputJ)
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

	body, err := fetchListValid(ctx)
	if err != nil {
		return err
	}

	for _, d := range body.Data {
		fmt.Println(fmt.Sprintf("%s = %s", d.Key, d.Value))
	}

	return nil
}
