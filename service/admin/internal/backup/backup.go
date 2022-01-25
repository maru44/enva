package backup

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

const buckUpMainDBFile = "./buckup/main_%s.sql"

func BackUp() error {
	cmd := exec.Command(
		"pg_dump",
		os.Getenv("POSTGRES_URL"),
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(fmt.Sprintf(buckUpMainDBFile, time.Now().Format("200601021504")), bytes, 0600); err != nil {
		fmt.Print(err)
	}

	return nil
}
