package commands

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	kvCreateBody struct {
		Id string `json:"data"`
	}
)

func fetchCreateKv(ctx context.Context, key, value string) (*kvCreateBody, error) {
	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	path := "/cli/kv/create?projectSlug=" + s.ProjectSlug
	if s.OrgSlug != nil {
		path += "&orgSlug=" + *s.OrgSlug
	}

	email, password, err := inputEmailPassword()
	if err != nil {
		return nil, err
	}

	input := domain.KvInputWithProjectID{
		Input: domain.KvInput{
			Key:   domain.KvKey(strings.Trim(key, "\"")),
			Value: domain.KvValue(value),
		},
	}
	inputJ, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := request(path, http.MethodPost, inputJ, email, password)
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
		body := &kvCreateBody{}
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
