package commands

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	initialize struct{}
)

func init() {
	Commands["init"] = func() domain.ICommandInteractor {
		return &initialize{}
	}
}

var (
	frontUrl = ""
)

func (c *initialize) Run(ctx context.Context, opts ...string) error {
	kvs, err := kvsFromEnvFile()
	if err != nil {
		return err
	}

	inputs := make([]domain.KvInput, len(kvs))
	count := 0
	for _, kv := range kvs {
		input := kv.ToInput()
		if err := input.Validate(); err != nil {
			return err
		}
		inputs[count] = *input
		count++
	}

	email, password, err := inputEmailPassword()
	if err != nil {
		return err
	}

	if _, err := fetchBulkInsertKvs(ctx, inputs, email, password); err != nil {
		if err.Error() == "Project is not found" {
			color.Yellow("Do you want to create a new project? (y/n):")
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			isCreate := scan.Text()
			if isCreate != "y" && isCreate != "Y" {
				return err
			}

			fmt.Print("description (can be blank): ")
			scan = bufio.NewScanner(os.Stdin)
			scan.Scan()
			desc := scan.Text()

			// @TODO post cli kv create
			// if success retry fetchBulkInsertKvs
			s, err := readSettings()
			if err != nil {
				return err
			}
			if _, err := fetchCreateProject(ctx, s.OrgSlug, s.ProjectSlug, desc, email, password); err != nil {
				return err
			}

			if _, err := fetchBulkInsertKvs(ctx, inputs, email, password); err != nil {
				return err
			}

			var path string
			if s.OrgSlug == nil {
				path = "project/" + s.ProjectSlug
			} else {
				path = "project/" + *s.OrgSlug + "/" + s.ProjectSlug
			}
			color.Green("Project created!\nURL: ")
			color.Blue("%s/%s", frontUrl, path)
			return nil
		}
		return err
	}

	color.Green("init project is succeded")
	return nil
}

func (c *initialize) Explain() string {
	return `	Setting key-value sets of remote based on local env file written in enva.json.
	If project written in enva.json does not exists, you can create project and set key-value sets of remote based on local env file. (This is optional. You can deny it.)

	* This command is so powerful that you can't execute if any remote key-value is set in the project.
`
}
