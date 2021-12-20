package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	kvCreateBody struct {
		Id string `json:"id"`
	}
)

func fetchCreateKv(ctx context.Context, key, value string) (*kvCreateBody, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/cli/kv/create?projectSlug=%s", ApiUrl, s.ProjectSlug)
	if s.OrgSlug != nil {
		// @TODO get by org
		// url =
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

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(inputJ),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", email+domain.CLI_HEADER_SEP+password)

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
