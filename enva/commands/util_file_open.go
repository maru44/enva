package commands

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/maru44/enva/service/api/pkg/domain"
)

func fileWriteFromResponse(body kvListBody) error {
	s, err := readSettings()
	if err != nil {
		return err
	}

	fileName := s.EnvFileName
	// if !strings.HasPrefix(fileName, ".") && !strings.HasPrefix(fileName, "/") && !strings.HasPrefix(fileName, "~") {
	// 	path, err := os.Getwd()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fileName = path + "/" + fileName
	// }

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

	if s.PreSentence != nil {
		if _, err := file.WriteString(*s.PreSentence + "\n"); err != nil {
			return err
		}
	}
	for _, d := range body.Data {
		if _, err := file.WriteString(fw(d)); err != nil {
			return err
		}
	}
	if s.SufSentence != nil {
		if _, err := file.WriteString(*s.SufSentence + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func fileReadAndUpdateKv(key, value string) error {
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

	kvs, err := kvsFromEnvFile()
	if err != nil {
		return err
	}
	// add or update
	kvs[domain.KvKey(key)] = domain.KvValue(value)

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if s.PreSentence != nil {
		if _, err := file.WriteString(*s.PreSentence + "\n"); err != nil {
			return err
		}
	}
	for k, v := range kvs {
		kv := domain.KvValid{
			Key:   k,
			Value: v,
		}
		if _, err := file.WriteString(fw(kv)); err != nil {
			return err
		}
	}
	if s.SufSentence != nil {
		if _, err := file.WriteString(*s.SufSentence + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func fileReadAndDeleteKv(key string) error {
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

	kvs, err := kvsFromEnvFile()
	if err != nil {
		return err
	}
	_, ok = kvs[domain.KvKey(key)]
	if ok {
		delete(kvs, domain.KvKey(key))
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if s.PreSentence != nil {
		if _, err := file.WriteString(*s.PreSentence + "\n"); err != nil {
			return err
		}
	}
	for k, v := range kvs {
		kv := domain.KvValid{
			Key:   k,
			Value: v,
		}
		if _, err := file.WriteString(fw(kv)); err != nil {
			return err
		}
	}
	if s.SufSentence != nil {
		if _, err := file.WriteString(*s.SufSentence + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func kvsFromEnvFile() (map[domain.KvKey]domain.KvValue, error) {
	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	fileName := s.EnvFileName

	ms, err := godotenv.Read(fileName)
	if err != nil {
		return nil, err
	}
	kvs := make(map[domain.KvKey]domain.KvValue, len(ms))
	for k, v := range ms {
		kvs[domain.KvKey(k)] = domain.KvValue(v)
	}
	return kvs, nil
}
