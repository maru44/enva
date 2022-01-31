package tag

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed tfvar.tmpl
var renderTemplate string

func tfgen(data *tagSt) error {
	t, err := template.New("image_tag_gen").Parse(renderTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to render with template: " + err.Error())
	}

	path := filepath.Join("./platform/terraform/prod", "gen_image_tag.tf")
	if err := os.WriteFile(path, buf.Bytes(), 0600); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
