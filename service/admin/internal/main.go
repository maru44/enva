package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/maru44/enva/enva/commands"
	"github.com/maru44/enva/service/admin/internal/backup"
	"github.com/maru44/enva/service/admin/internal/privacy"
)

type (
	versionOs struct {
		Name  string   `json:"os"`
		Archs []string `json:"archs"`
	}

	version struct {
		Version   string      `json:"version"`
		Oss       []versionOs `json:"oss"`
		UpdatedAt string      `json:"updated_at"`
	}

	explain struct {
		Name    string `json:"command"`
		Explain string `json:"explain"`
	}
)

const (
	tarFile     = "./service/front/public/enva/tar.json"
	explainFile = "./service/front/src/components/cli/explain.json"
)

func main() {
	if len(commands.Commands) != len(commands.AllCommands) {
		panic("commands length not correspond\ncommands.Commands with commands.AllCommands")
	}

	flag.Parse()
	args := flag.Args()

	if args[0] == "tar/json" {
		// if env is local, skip gen tar.json
		if os.Getenv("CLI_API_URL") == "http://localhost:8080" {
			fmt.Println("skip to overwrite tar.json")
			return
		}

		if len(args) != 4 {
			panic("invalid args: need 3 args")
		}

		updateFrontVersionFile(args[1], args[2], args[3])
		return
	}

	if args[0] == "explain/json" {
		overwriteExplainFile()
		return
	}

	if args[0] == "privacy/json" {
		if err := privacy.GenPrivacyJson(); err != nil {
			panic(err)
		}
		return
	}

	if args[0] == "backup" {
		if err := backup.BackUp(); err != nil {
			panic(err)
		}
		fmt.Println("succeeded to backup db!")
		return
	}

	panic("no such commands")
}

func overwriteExplainFile() {
	explains := make([]explain, len(commands.AllCommands))
	for i, name := range commands.AllCommands {
		cmd := commands.Commands[name]
		exp := explain{
			Name:    name,
			Explain: cmd().Explain(),
		}
		explains[i] = exp
	}

	j, err := json.Marshal(explains)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(explainFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteAt(j, 0); err != nil {
		panic(err)
	}
}

func updateFrontVersionFile(inputVersion, inputOs, inputArch string) {
	var (
		vs                []version
		idxVersion, idxOs int = -1, -1
	)
	data, err := ioutil.ReadFile(tarFile)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &vs); err != nil && len(data) != 0 {
		panic(err)
	}
	for i, v := range vs {
		if v.Version == inputVersion {
			idxVersion = i
			break
		}
	}

	// if input version does not ex
	if idxVersion == -1 {
		newValue := version{
			Version: inputVersion,
			Oss: []versionOs{
				{
					Name: inputOs,
					Archs: []string{
						inputArch,
					},
				},
			},
			UpdatedAt: time.Now().Format("Jan 2, 2006"),
		}
		vs = append(vs, newValue)
	}

	// if input version ex
	if idxVersion != -1 {
		for i, o := range vs[idxVersion].Oss {
			if o.Name == inputOs {
				idxOs = i
				break
			}
		}

		// if input os does not ex
		if idxOs == -1 {
			vs[idxVersion].Oss = append(vs[idxVersion].Oss, versionOs{Name: inputOs, Archs: []string{inputArch}})
		} else {
			// if input os ex
			for _, a := range vs[idxVersion].Oss[idxOs].Archs {
				// if input arch ex
				if a == inputArch {
					return
				}
			}
			// if arch not ex
			vs[idxVersion].Oss[idxOs].Archs = append(vs[idxVersion].Oss[idxOs].Archs, inputArch)
		}
	}

	j, err := json.Marshal(vs)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(tarFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteAt(j, 0); err != nil {
		panic(err)
	}
}
