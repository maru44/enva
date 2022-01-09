package commands

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

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

	projectSlug, orgSlug, fileName, pre, suf, err := c.inputSettingsInfo()
	if err != nil {
		return err
	}

	file, err := os.OpenFile("enva.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(c.fileStr(projectSlug, orgSlug, fileName, pre, suf))
	fmt.Println("\nSucceeded to create enva.json")

	return nil
}

func (c *set) Explain() string {
	return `
	create enva.json (only if enva.json does not exists)
`
}

func (c *set) inputSettingsInfo() (projectSlug, orgSlug, fileName, pre, suf string, err error) {
	fmt.Println("start creating enva.json ...")
	fmt.Print("project slug: ")
	for {
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		projectSlug = scan.Text()

		if projectSlug != "" {
			// orgSlug
			fmt.Println("org slug (can be blank): ")
			scan := bufio.NewScanner(os.Stdin)
			scan.Scan()
			orgSlug = scan.Text()

			// filePath
			fmt.Print("filepath: ")
			scan = bufio.NewScanner(os.Stdin)
			scan.Scan()
			fileName = scan.Text()

			if fileName == "" {
				return "", "", "", "", "", errors.New("filepath must not be blank")
			}

			fmt.Print("pre sentence (can be blank): ")
			scan = bufio.NewScanner(os.Stdin)
			scan.Scan()
			pre = scan.Text()

			// filePath
			fmt.Print("suf sentence (can be blank): ")
			scan = bufio.NewScanner(os.Stdin)
			scan.Scan()
			suf = scan.Text()

			return
		}
		return "", "", "", "", "", errors.New("project slug must not be blank")
	}
}

func (c *set) fileStr(projectSlug, orgSlug, fileName, pre, suf string) string {
	input := []string{
		fmt.Sprintf(`	"env_file_name": "%s"`, fileName),
		fmt.Sprintf(`	"project_slug": "%s"`, projectSlug),
	}

	if orgSlug != "" {
		input = append(input, fmt.Sprintf(`	"org_slug": "%s"`, orgSlug))
	}
	if pre != "" {
		input = append(input, fmt.Sprintf(`	"pre_sentence": "%s"`, pre))
	}
	if suf != "" {
		input = append(input, fmt.Sprintf(`	"suf_sentence": "%s"`, suf))
	}

	jsonContent := strings.Join(input, ",\n")

	return fmt.Sprintf("{\n%s\n}\n", jsonContent)
}
