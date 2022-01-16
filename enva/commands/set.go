package commands

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
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
	if _, err := readSettings(); err == nil {
		return errors.New("enva.json already exists")
	}

	setting, err := c.inputSettingsInfo()
	if err != nil {
		return err
	}

	file, err := os.OpenFile("enva.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("{\n" + *setting + "\n}\n"); err != nil {
		return err
	}
	color.Green("\nSucceeded to create enva.json")

	return nil
}

func (c *set) Explain() string {
	return `	create enva.json (only if enva.json does not exists in current directory)
`
}

func (c *set) inputSettingsInfo() (*string, error) {
	fmt.Println("start creating enva.json ...")
	projectSlug, err := c.input("project slug", true)
	if err != nil {
		return nil, err
	}

	orgSlug, err := c.input("org slug", false)
	if err != nil {
		return nil, err
	}

	fileName, err := c.input("filepath", true)
	if err != nil {
		return nil, err
	}

	pre, err := c.input("pre sentence", false)
	if err != nil {
		return nil, err
	}

	suf, err := c.input("suf sentence", false)
	if err != nil {
		return nil, err
	}

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

	return &jsonContent, nil
}

func (c *set) input(field string, isRequired bool) (string, error) {
	if isRequired {
		fmt.Print(field, " *: ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		ret := scan.Text()
		if ret == "" {
			return "", errors.New(field + " must not be blank")
		}
		return ret, nil
	}

	fmt.Print(field, ": ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	return scan.Text(), nil
}
