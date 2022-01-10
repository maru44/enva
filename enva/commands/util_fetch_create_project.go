package commands

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/tools"
)

type (
	projectCreateBody struct {
		ID string `json:"data"`
	}
)

func fetchCreateProject(ctx context.Context, desc, email, password string) (*projectCreateBody, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	path := "/cli/project/create"
	input := domain.CliProjectInput{
		Name:        s.ProjectSlug,
		Slug:        s.ProjectSlug,
		Description: tools.StringPtr(desc),
		OrgSlug:     s.OrgSlug,
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
		body := &projectCreateBody{}
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
