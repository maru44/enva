package main

import (
	"context"
	"flag"

	"github.com/maru44/enva/enva/commands"
)

func main() {
	ctx := context.Background()
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		commands.Run(ctx, "version")
		commands.Run(ctx, "help")
		return
	}

	name := args[0]
	opts := args[1:]

	commands.Run(ctx, name, opts...)
}
