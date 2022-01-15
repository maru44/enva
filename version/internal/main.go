package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	versionOs struct {
		Name  string   `json:"os"`
		Archs []string `json:"archs"`
	}

	version struct {
		Version string      `json:"version"`
		Oss     []versionOs `json:"oss"`
	}
)

const fileName = "./service/front/public/enva/tar.json"

func main() {
	// if env is local skip
	if os.Getenv("CLI_API_URL") == "http://localhost:8080" {
		fmt.Println("skip to overwrite tar.json")
		return
	}

	flag.Parse()
	args := flag.Args()

	if len(args) != 3 {
		panic("invalid args: need 3 args")
	}
	inputVersion := args[0]
	inputOs := args[1]
	inputArch := args[2]

	var (
		vs                []version
		idxVersion, idxOs int = -1, -1
	)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &vs); err != nil {
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

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteAt(j, 0); err != nil {
		panic(err)
	}
}
