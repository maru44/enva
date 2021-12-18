package main

import (
	"context"
	"flag"

	"github.com/maru44/enva/service/api/internal/infra"
	"github.com/maru44/enva/service/api/internal/interface/commands"
)

func main() {
	ctx := context.Background()

	sql := infra.NewSqlHandler()
	pass := &infra.Password{}

	flag.Parse()
	args := flag.Args()

	name := args[0]
	opts := args[1:]

	commands.Run(ctx, sql, pass, name, opts...)
}
