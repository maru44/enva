package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type (
	kvDeleteBody struct {
		Affected int `json:"data"`
	}
)

func fetchDeleteKv(ctx context.Context, key string) (*kvDeleteBody, error) {
	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/cli/kv/delete?projectSlug=%s&key=%s", s.ProjectSlug, key)
	if s.OrgSlug != nil {
		path += "&orgSlug=" + *s.OrgSlug
	}

	email, password, err := inputEmailPassword()
	if err != nil {
		return nil, err
	}

	req, err := request(path, http.MethodDelete, nil, email, password)
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
