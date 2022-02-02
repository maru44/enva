package commands

import (
	"encoding/json"
	"io/ioutil"

	"github.com/maru44/enva/service/api/pkg/domain"
)

// for read file
func kvsFromJsonFile(filename string) ([]domain.KvValid, error) {
	dataBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data map[string]string
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, err
	}

	var kvs []domain.KvValid
	for key, value := range data {
		kvs = append(kvs, domain.KvValid{
			Key:   domain.KvKey(key),
			Value: domain.KvValue(value),
		})
	}
	return kvs, nil
}
