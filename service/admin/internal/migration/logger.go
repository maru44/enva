package migration

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func newLogger(ctx context.Context) migrate.Logger {
	return &logger{}
}

type logger struct{}

func (l *logger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (l *logger) Verbose() bool {
	return false
}
