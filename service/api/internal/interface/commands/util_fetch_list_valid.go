package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/maru44/enva/service/api/internal/config"
)

func fetchListValid(ctx context.Context) (*kvListBody, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	var url string
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
	}
	url = fmt.Sprintf("%s/cli/kv?projectSlug=%s", config.API_URL, s.ProjectSlug)

	input, err := inputEmailPassword()
	if err != nil {
		return nil, err
	}

	inputJ, err := json.Marshal(input)
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
