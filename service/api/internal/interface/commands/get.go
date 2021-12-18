package commands

import (
	"context"
	"fmt"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	kvList struct{}

	errorBody struct {
		Error string `json:"error"`
	}

	kvListBody struct {
		Data []domain.KvValid `json:"data"`
	}
)

func init() {
	Commands["get"] = func() domain.ICommandInteractor {
		return &kvList{}
	}
}

func (c *kvList) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	body, err := fetchListValid(ctx)
	if err != nil {
		return err
	}

	for _, d := range body.Data {
		fmt.Println(fmt.Sprintf("%s=%s", d.Key, d.Value))
	}

	return nil
}
