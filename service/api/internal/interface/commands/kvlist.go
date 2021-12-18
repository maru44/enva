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
	kvList struct{}

	resBody struct {
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

	s, err := readSettings()
	if err != nil {
		return err
	}

	var url string
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
	}
	url = fmt.Sprintf("%s/cli/kv?projectSlug=%s", config.API_URL, s.ProjectSlug)

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

		for _, d := range body.Data {
			fmt.Println(fmt.Sprintf("%s=%s", d.Key, d.Value))
		}
	}

	return nil
}
