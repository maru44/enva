package commands

import (
	"fmt"
	"strings"

	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/tools"
)

var (
	fileWriteMap = map[string]func(kv domain.KvValid) string{
		".envrc":  writeDirenv,
		".tfvars": writeTfval,
		".yaml":   writeYaml,
		".yml":    writeYaml,
	}
)

func writeNormal(kv domain.KvValid) string {
	value := soroundedQuotes(kv.Value.String())
	return fmt.Sprintf("%s=%s\n", kv.Key, value)
}

func writeDirenv(kv domain.KvValid) string {
	value := soroundedQuotes(kv.Value.String())
	return fmt.Sprintf("export %s=%s\n", kv.Key, value)
}

func writeTfval(kv domain.KvValid) string {
	value := escapeDoubleQuotes(kv.Value.String())
	return fmt.Sprintf("%s=\"%s\"\n", kv.Key, value)
}

func writeYaml(kv domain.KvValid) string {
	value := soroundedQuotes(kv.Value.String())
	return fmt.Sprintf("%s: %s\n", kv.Key, value)
}

func writeJson(kv domain.KvValid, isLast bool) string {
	value := escapeDoubleQuotes(kv.Value.String())
	if isLast {
		return fmt.Sprintf("\t\"%s\": \"%s\"\n", kv.Key, value)
	}
	return fmt.Sprintf("\t\"%s\": \"%s\",\n", kv.Key, value)
}

/* utils */

func escapeDoubleQuotes(str string) string {
	return strings.ReplaceAll(str, "\"", "\\\"")
}

func soroundedQuotes(str string) string {
	isAllSlug := true
	for _, r := range str {
		index := strings.IndexRune(tools.SlugLetters, r)
		if index == -1 {
			isAllSlug = false
			break
		}
	}

	if !isAllSlug {
		return fmt.Sprintf("\"%s\"", escapeDoubleQuotes(str))
	}

	return str
}
