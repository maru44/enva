package tag

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type tagSt struct {
	ApiImage       int `json:"apiImageTag"`
	MigrationImage int `json:"migrationImageTag"`
}

const (
	tagFile = "./infra/docker/tag.json"
)

func SuccessApi() {
	t := read()

	// tf generate
	if err := tfgen(t); err != nil {
		panic(err)
	}

	// then json update
	t.ApiImage++
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	file := readToWrite()
	defer file.Close()

	if _, err := file.WriteAt(j, 0); err != nil {
		panic(err)
	}
}

func SuccessMigration() {
	t := read()

	// tf generate
	if err := tfgen(t); err != nil {
		panic(err)
	}

	// then json update
	t.MigrationImage++
	j, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	file := readToWrite()
	defer file.Close()

	if _, err := file.WriteAt(j, 0); err != nil {
		panic(err)
	}
}

func read() *tagSt {
	var tag *tagSt
	data, err := ioutil.ReadFile(tagFile)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &tag); err != nil && len(data) != 0 {
		panic(err)
	}
	return tag
}

func readToWrite() *os.File {
	file, err := os.OpenFile(tagFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	return file
}
