package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/maru44/enva/service/api/pkg/domain"
)

var (
	fileTypeMap = map[string]func(kv domain.KvValid) string{
		".envrc":  outputDirenv,
		".tfvars": outputTfval,
	}
)

func fileOpen() (*os.File, func(kv domain.KvValid) string, error) {
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
	ft, ok := fileTypeMap[ext]
	if !ok {
		ft = outputNormal
	}

	return file, ft, nil
}
