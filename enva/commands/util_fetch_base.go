package commands

import (
	"bytes"
	"net/http"

	"github.com/maru44/enva/service/api/pkg/domain"
)

func request(path string, method string, jsonInput []byte, email, password string) (*http.Request, error) {
	req, err := http.NewRequest(method, apiUrl+path, bytes.NewBuffer(jsonInput))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", email+domain.CLI_HEADER_SEP+password)

	return req, nil
}
