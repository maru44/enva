package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/maru44/enva/service/api/internal/config"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	kvBulkInsertBody struct {
		Data string `json:"data"`
	}
)

func fetchBulkInsertKvs(ctx context.Context, inputs []domain.KvInput) (*kvBulkInsertBody, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/cli/kv/create/bulk?projectSlug=%s", config.API_URL, s.ProjectSlug)
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
	}

	email, password, err := inputEmailPassword()
	if err != nil {
		return nil, err
	}

	inputJ, err := json.Marshal(inputs)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(inputJ),
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
		body := &kvBulkInsertBody{}
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
