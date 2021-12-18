package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maru44/enva/service/api/internal/config"
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

	s, err := readSettings()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/cli/kv?projectSlug=%s", config.API_URL, s.ProjectSlug)
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
	}

	input, err := inputEmailPassword()
	if err != nil {
		return err
	}

	inputJ, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(inputJ),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var body resBody
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return err
		}

		if err := overWriteEnv(body); err != nil {
			return err
		}
	default:
	}

	return nil
}

func overWriteEnv(body resBody) error {
	file, outputFunc, err := fileOpen()
	if err != nil {
		return err
	}
	defer file.Close()

	for _, d := range body.Data {
		if _, err := file.WriteString(outputFunc(d)); err != nil {
			return err
		}
	}
	return nil
}
