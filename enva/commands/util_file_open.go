package commands

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/maru44/enva/service/api/pkg/domain"
)

func fileWriteFromResponse(body kvListBody) error {
	s, err := readSettings()
	if err != nil {
		return err
	}

	fileName := s.EnvFileName
	if !strings.HasPrefix(fileName, ".") && !strings.HasPrefix(fileName, "/") && !strings.HasPrefix(fileName, "~") {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		fileName = path + "/" + fileName
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileWriteMap[ext]
	if !ok {
		f = writeNormal
	}

	if s.PreSentence != nil {
		file.WriteString(*s.PreSentence + "\n")
	}
	for _, d := range body.Data {
		if _, err := file.WriteString(f(d)); err != nil {
			return err
		}
	}
	if s.SufSentence != nil {
		file.WriteString(*s.SufSentence + "\n")
	}
	return nil
}

func fileReadAndUpdateKv(key, value string) error {
	s, err := readSettings()
	if err != nil {
		return err
	}

	fileName := s.EnvFileName
	if !strings.HasPrefix(fileName, ".") && !strings.HasPrefix(fileName, "/") && !strings.HasPrefix(fileName, "~") {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		fileName = path + "/" + fileName
	}

	// to read
	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileReadMap[ext]
	if !ok {
		f = readOneLineNormal
	}

	fileRead, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
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

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	/* write file by created kvs */
	fw, ok := fileWriteMap[ext]
	if !ok {
		fw = writeNormal
	}

	if s.PreSentence != nil {
		file.WriteString(*s.PreSentence + "\n")
	}
	for _, d := range kvs {
		if _, err := file.WriteString(fw(d)); err != nil {
			return err
		}
	}
	if s.SufSentence != nil {
		file.WriteString(*s.SufSentence + "\n")
	}
	return nil
}

func fileReadAndDeleteKv(key string) error {
	s, err := readSettings()
	if err != nil {
		return err
	}

	fileName := s.EnvFileName
	if !strings.HasPrefix(fileName, ".") && !strings.HasPrefix(fileName, "/") && !strings.HasPrefix(fileName, "~") {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		fileName = path + "/" + fileName
	}

	// to read
	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileReadMap[ext]
	if !ok {
		f = readOneLineNormal
	}

	fileRead, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer fileRead.Close()

	var kvs []domain.KvValid

	// crete kvs from file
	scanner := bufio.NewScanner(fileRead)
	for scanner.Scan() {
		kv := f(scanner.Text())
		if kv != nil {
			if kv.Key != domain.KvKey(key) {
				kvs = append(kvs, *kv)
			}
		}
	}

	fileRead.Close()

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	/* write file by created kvs */
	fw, ok := fileWriteMap[ext]
	if !ok {
		fw = writeNormal
	}

	if s.PreSentence != nil {
		file.WriteString(*s.PreSentence + "\n")
	}
	for _, d := range kvs {
		if _, err := file.WriteString(fw(d)); err != nil {
			return err
		}
	}
	if s.SufSentence != nil {
		file.WriteString(*s.SufSentence + "\n")
	}
	return nil
}

func fileReadAndCreateKvs() ([]domain.KvValid, error) {
	s, err := readSettings()
	if err != nil {
		return nil, err
	}

	fileName := s.EnvFileName
	if !strings.HasPrefix(fileName, ".") && !strings.HasPrefix(fileName, "/") && !strings.HasPrefix(fileName, "~") {
		path, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		fileName = path + "/" + fileName
	}

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ext := filepath.Ext(s.EnvFileName)
	f, ok := fileReadMap[ext]
	if !ok {
		f = readOneLineNormal
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
