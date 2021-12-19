package commands

import (
	"fmt"
	"strings"

	"github.com/maru44/enva/service/api/pkg/domain"
)

// for read file

var (
	fileInputMap = map[string]func(string) *domain.KvValid{
		".envrc": inputDirenv,
	}
)

func inputNormal(str string) *domain.KvValid {
	if str == "" {
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

func inputDirenv(str string) *domain.KvValid {
	if str == "" {
		return nil
	}

	sp := splitEqual(str)
	if len(sp) != 2 {
		return nil
	}

	fmt.Println("raw", sp[1])
	fmt.Println("str", string(sp[1]))

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

func splitEqual(str string) []string {
	return strings.SplitN(str, "=", 2)
}
