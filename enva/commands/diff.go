package commands

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type (
	diff struct{}

	diffKvValid struct {
		Key         domain.KvKey
		RemoteValue domain.KvValue
		LocalValue  domain.KvValue
	}
)

func init() {
	Commands["diff"] = func() domain.ICommandInteractor {
		return &diff{}
	}
}

func (c *diff) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	body, err := fetchListValid(ctx)
	if err != nil {
		return err
	}

	remoteData := domain.MapFromKv(body.Data)

	localDataRaw, err := fileReadAndCreateKvs()
	localData := domain.MapFromKv(localDataRaw)
	if err != nil {
		return err
	}

	var (
		diffs                         []diffKvValid
		listOnlyLocal, listOnlyRemote []domain.KvValid
	)

	for k, v := range remoteData {
		lv, ok := localData[k]
		if !ok {
			listOnlyRemote = append(listOnlyRemote, domain.KvValid{Key: k, Value: v})
			continue
		}

		if lv == v {
			continue
		}

		diffs = append(diffs, diffKvValid{
			Key:         k,
			RemoteValue: v,
			LocalValue:  lv,
		})
	}

	for k, v := range localData {
		if _, ok := remoteData[k]; !ok {
			listOnlyLocal = append(listOnlyLocal, domain.KvValid{Key: k, Value: v})
		}
	}

	allDiffs := len(diffs) + len(listOnlyLocal) + len(listOnlyRemote)
	if allDiffs == 0 {
		fmt.Print("There are no difference between remote and local!\n\n")
		return nil
	}

	if diffs != nil {
		color.HiGreen("There are %d differences.", len(diffs))
		for _, d := range diffs {
			fmt.Printf("%s:\n\tremote: %s\n\tlocal: %s\n", d.Key, d.RemoteValue, d.LocalValue)
		}
		fmt.Print("\n")
	}

	if listOnlyLocal != nil {
		color.HiMagenta("Local Only: ")
		for _, kv := range listOnlyLocal {
			fmt.Printf("%s (%s)\n", kv.Key, kv.Value)
		}
		fmt.Print("\n")
	}

	if listOnlyRemote != nil {
		color.HiCyan("Remote Only: ")
		for _, kv := range listOnlyRemote {
			fmt.Printf("%s (%s)\n", kv.Key, kv.Value)
		}
		fmt.Print("\n")
	}

	return nil
}

func (c *diff) Explain() string {
	return `	Getting the difference between remote and local key-value sets and output them in command line.
`
}
