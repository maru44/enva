package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/maru44/enva/service/api/internal/config"
)

type (
	kvDeleteBody struct {
		Affected int `json:"data"`
	}
)

func fetchDeleteKv(ctx context.Context, key string) (*kvDeleteBody, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/cli/kv/delete?projectSlug=%s&key=%s", config.API_URL, s.ProjectSlug, key)
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
	}

	email, password, err := inputEmailPassword()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", email+config.CLI_HEADER_SEP+password)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		body := &kvDeleteBody{}
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return nil, err
		}

		return body, nil
	default:
		body := &errorBody{}
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			return nil, err
		}

		return nil, errors.New(body.Error)
	}
}
