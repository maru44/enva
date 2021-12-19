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
