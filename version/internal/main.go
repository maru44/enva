package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const fileName = "./service/front/public/enva/version.json"

func main() {
	flag.Parse()
	args := flag.Args()

	version := args[0]
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fi, _ := file.Stat()
	leng := fi.Size()
	if leng == 0 {
		if _, err := file.Write([]byte(fmt.Sprintf(`["%s"]`, version))); err != nil {
			panic(err)
		}
	} else {
		var versions []string
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(data, &versions); err != nil {
			panic(err)
		}
		for _, v := range versions {
			if v == version {
				return
			}
		}

		if _, err := file.WriteAt([]byte(fmt.Sprintf(`,"%s"]`, version)), leng-1); err != nil {
			panic(err)
		}
	}
}
