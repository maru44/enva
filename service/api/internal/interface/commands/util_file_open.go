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

func fileOpenToWrite() (*os.File, func(kv domain.KvValid) string, error) {
	s, err := readSettings()
	if err != nil {
		return nil, nil, err
	}

	path, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/%s", path, s.EnvFileName), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileOutputMap[ext]
	if !ok {
		f = outputNormal
	}

	return file, f, nil
}

func fileReadToKvs() ([]domain.KvValid, error) {
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
