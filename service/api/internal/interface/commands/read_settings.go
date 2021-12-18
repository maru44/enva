package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/maru44/enva/service/api/pkg/domain"
)

func readSettings() (*domain.Settings, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("%s/enva.json", path))
	if err != nil {
		return nil, err
	}

	settings := &domain.Settings{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

	return settings, nil
}
