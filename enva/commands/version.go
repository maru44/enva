package commands

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed version.json
var versionBytes []byte

type (
	version struct{}

	versionInfo struct {
		Version  string `json:"version"`
		Updation string `json:"updation"`
	}
)

func (c *version) Run(ctx context.Context, opts ...string) error {
	var info versionInfo
	if err := json.Unmarshal(versionBytes, &info); err != nil {
		return err
	}
	fmt.Printf(`Enva!!

Version: %s
Updation: %s

`, info.Version, info.Updation)
	return nil
}

func (c *version) Explain() string {
	return `
	Showing version and information of updation.
`
}
