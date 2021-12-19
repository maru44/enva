package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/maru44/enva/service/api/pkg/domain"
)

var (
	fileOutputMap = map[string]func(kv domain.KvValid) string{
		".envrc":  outputDirenv,
		".tfvars": outputTfval,
	}
)

func fileWriteFromResponse(body kvListBody) error {
	s, err := readSettings()
	if err != nil {
		return err
	}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", path, s.EnvFileName), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileOutputMap[ext]
	if !ok {
		f = outputNormal
	}

	for _, d := range body.Data {
		if _, err := file.WriteString(f(d)); err != nil {
			return err
		}
	}
	return nil
}

func fileWriteToUpdateKv(key, value string) error {
	s, err := readSettings()
	if err != nil {
		return err
	}

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	// to read
	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileInputMap[ext]
	if !ok {
		f = inputNormal
	}

	fileRead, err := os.OpenFile(fmt.Sprintf("%s/%s", path, s.EnvFileName), os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer fileRead.Close()

	var (
		kvs             []domain.KvValid
		isExistsAlready bool
	)
	// crete kvs from file
	scanner := bufio.NewScanner(fileRead)
	for scanner.Scan() {
		kv := f(scanner.Text())
		if kv != nil {
			if kv.Key == domain.KvKey(key) {
				kvs = append(kvs, domain.KvValid{
					Key:   kv.Key,
					Value: domain.KvValue(value),
				})
				isExistsAlready = true
				continue
			}
			kvs = append(kvs, *kv)
		}
	}

	// if new key
	if !isExistsAlready {
		kvs = append(kvs, domain.KvValid{
			Key:   domain.KvKey(key),
			Value: domain.KvValue(value),
		})
	}
	fileRead.Close()

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", path, s.EnvFileName), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	/* write file by created kvs */
	fw, ok := fileOutputMap[ext]
	if !ok {
		fw = outputNormal
	}

	for _, d := range kvs {
		if _, err := file.WriteString(fw(d)); err != nil {
			return err
		}
	}
	return nil
}

func fileReadAndCreateKvs() ([]domain.KvValid, error) {
	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", path, s.EnvFileName), os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileInputMap[ext]
	if !ok {
		f = inputNormal
	}

	scanner := bufio.NewScanner(file)

	var kvs []domain.KvValid
	for scanner.Scan() {
		kv := f(scanner.Text())
		if kv != nil {
			kvs = append(kvs, *kv)
		}
	}

	return kvs, nil
}
