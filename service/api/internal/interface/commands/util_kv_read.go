package commands

import (
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

	removedR := strings.TrimRight(sp[1], "\n")

	return &domain.KvValid{
		Key:   domain.KvKey(sp[0]),
		Value: domain.KvValue(strings.Trim(removedR, "\"")),
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

	key := strings.TrimLeft(sp[0], "export ")
	val := strings.TrimRight(sp[1], "\n")

	return &domain.KvValid{
		Key:   domain.KvKey(key),
		Value: domain.KvValue(strings.Trim(val, "\"")),
	}
}

func splitEqual(str string) []string {
	return strings.SplitN(str, "=", 2)
}
