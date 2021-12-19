package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/maru44/enva/service/api/internal/config"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	errorBody struct {
		Error string `json:"error"`
	}

	kvListBody struct {
		Data []domain.KvValid `json:"data"`
	}
)

func fetchListValid(ctx context.Context) (*kvListBody, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/cli/kv?projectSlug=%s", config.API_URL, s.ProjectSlug)
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
	}

	email, password, err := inputEmailPassword()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodGet,
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
		body := &kvListBody{}
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
