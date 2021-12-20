package commands

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	set struct{}
)

func init() {
	Commands["set"] = func() domain.ICommandInteractor {
		return &set{}
	}
}

func (c *set) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if _, err := readSettings(); err == nil {
		return errors.New("enva.json already exists")
	}

	projectSlug, fileName, _, err := c.inputSettingsInfo()
	if err != nil {
		return err
	}

	file, err := os.OpenFile("enva.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(
		fmt.Sprintf(
			`{
	"project_slug": "%s",
	"env_file_name": "%s"
}
`,
			projectSlug,
			fileName,
		),
	)

	fmt.Println("\nSucceeded to create enva.json")

	return nil
}

func (c *set) Explain() string {
	return `
	create enva.json (only if enva.json does not exists)
`
}

func (c *set) inputSettingsInfo() (projectSlug string, fileName string, orgId string, err error) {
	fmt.Println("start creating enva.json ...")
	fmt.Print("project slug: ")
	for {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		projectSlug = scan.Text()

		if projectSlug != "" {
			fmt.Print("filepath: ")
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			fileName = scan.Text()

			if fileName == "" {
				return "", "", "", errors.New("filepath must not be blank")
			}

			return
		}
		return "", "", "", errors.New("project slug must not be blank")
	}
}
