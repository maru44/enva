package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	kvDetailBody struct {
		Data domain.KvValid `json:"data"`
	}
)

func fetchDetailValid(ctx context.Context, key, email, password string) (*kvDetailBody, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/cli/kv/detail?key=%s&projectSlug=%s", key, s.ProjectSlug)
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
	}

	req, err := request(path, http.MethodGet, nil, email, password)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		body := &kvDetailBody{}
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
