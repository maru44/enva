package commands

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/joho/godotenv"
	"github.com/maru44/enva/service/api/pkg/domain"
)

func fileWriteFromResponse(body kvListBody) error {
	return writeKvsToFile(body.Data)
}

func fileReadAndUpdateKv(key, value string) error {
	kvs, err := kvsFromEnvFile()
	if err != nil {
		return err
	}

	exists := false
	for i, kv := range kvs {
		if kv.Key == domain.KvKey(key) {
			kvs[i].Value = domain.KvValue(value)
			exists = true
		}
	}
	if !exists {
		kvs = append(kvs, domain.KvValid{Key: domain.KvKey(key), Value: domain.KvValue(value)})
	}

	return writeKvsToFile(kvs)
}

func fileReadAndDeleteKv(key string) error {
	kvs, err := kvsFromEnvFile()
	if err != nil {
		return err
	}
	for i, kv := range kvs {
		if kv.Key == domain.KvKey(key) {
			kvs = kvs[:i+copy(kvs[i:], kvs[i+1:])]
		}
	}

	return writeKvsToFile(kvs)
}

/*******************************
	utils
*******************************/

func writeKvsToFile(kvs []domain.KvValid) error {
	s, err := readSettings()
	if err != nil {
		return err
	}

	fileName := s.EnvFileName
	// to read
	ext := filepath.Ext(s.EnvFileName)

	/* write file by created kvs */
	fw, ok := fileWriteMap[ext]
	if !ok {
		fw = writeNormal
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	sort.Slice(kvs, func(i, j int) bool { return kvs[i].Key.String() > kvs[j].Key.String() })
	if ext == ".json" {
		if _, err := file.WriteString("{" + "\n"); err != nil {
			return err
		}
		length := len(kvs)
		for i, kv := range kvs {
			if _, err := file.WriteString(writeJson(kv, length == i+1)); err != nil {
				return err
			}
		}
		if _, err := file.WriteString("}" + "\n"); err != nil {
			return err
		}
		return nil
	}

	// pre
	if s.PreSentence != nil {
		if _, err := file.WriteString(*s.PreSentence + "\n"); err != nil {
			return err
		}
	}
	// kvs
	for _, kv := range kvs {
		if _, err := file.WriteString(fw(kv)); err != nil {
			return err
		}
	}
	// suf
	if s.SufSentence != nil {
		if _, err := file.WriteString(*s.SufSentence + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func kvsFromEnvFile() ([]domain.KvValid, error) {
	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	fileName := s.EnvFileName

	ext := filepath.Ext(fileName)
	// json
	if ext == ".json" {
		kvs, err := kvsFromJsonFile(fileName)
		if err != nil {
			return nil, err
		}
		return kvs, nil
	}

	// others
	ms, err := godotenv.Read(fileName)
	if err != nil {
		return nil, err
	}
	kvs := make([]domain.KvValid, len(ms))
	count := 0
	for k, v := range ms {
		kvs[count] = domain.KvValid{
			Key:   domain.KvKey(k),
			Value: domain.KvValue(v),
		}
		count++
	}
	return kvs, nil
}
