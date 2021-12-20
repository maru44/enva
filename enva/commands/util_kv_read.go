package commands

import (
	"strings"

	"github.com/maru44/enva/service/api/pkg/domain"
)

// for read file

var (
	fileReadMap = map[string]func(string) *domain.KvValid{
		".envrc":  readOneLineDirenv,
		".tfvars": readOneLineTfvals,
	}
)

func readOneLineNormal(str string) *domain.KvValid {
	if str == "" {
		return nil
	}

	if strings.HasPrefix(str, "#") {
		return nil
	}

	sp := splitEqual(str)
	if len(sp) != 2 {
		return nil
	}

	removedR := strings.TrimRight(string(sp[1]), "\n")
	trimed := strings.Trim(removedR, "\"")

	val := strings.ReplaceAll(trimed, "\\\\", "バックスラッシュ")
	val = strings.ReplaceAll(val, "\\", "")
	val = strings.ReplaceAll(val, "バックスラッシュ", "\\\\")

	return &domain.KvValid{
		Key:   domain.KvKey(sp[0]),
		Value: domain.KvValue(val),
	}
}

func readOneLineDirenv(str string) *domain.KvValid {
	if str == "" {
		return nil
	}

	if strings.HasPrefix(str, "#") {
		return nil
	}

	sp := splitEqual(str)
	if len(sp) != 2 {
		return nil
	}

	key := strings.TrimLeft(sp[0], "export ")
	removeR := strings.TrimRight(sp[1], "\n")
	trimed := strings.Trim(removeR, "\"")

	val := strings.ReplaceAll(trimed, "\\\\", "バックスラッシュ")
	val = strings.ReplaceAll(val, "\\", "")
	val = strings.ReplaceAll(val, "バックスラッシュ", "\\\\")

	return &domain.KvValid{
		Key:   domain.KvKey(key),
		Value: domain.KvValue(val),
	}
}

func readOneLineTfvals(str string) *domain.KvValid {
	if str == "" {
		return nil
	}

	if strings.HasPrefix(str, "#") || strings.HasPrefix(str, "//") {
		return nil
	}

	sp := splitEqual(str)
	if len(sp) != 2 {
		return nil
	}

	removedR := strings.TrimRight(string(sp[1]), "\n")
	trimed := strings.Trim(removedR, "\"")

	val := strings.ReplaceAll(trimed, "\\\\", "バックスラッシュ")
	val = strings.ReplaceAll(val, "\\", "")
	val = strings.ReplaceAll(val, "バックスラッシュ", "\\\\")

	return &domain.KvValid{
		Key:   domain.KvKey(sp[0]),
		Value: domain.KvValue(val),
	}
}

func splitEqual(str string) []string {
	return strings.SplitN(str, "=", 2)
}
