package database

import (
	"context"
	"database/sql"
)

type (
	SqlHandlerAbstract interface {
		QueryContext(context.Context, string, ...interface{}) (RowsAbstract, error)
		QueryRowContext(context.Context, string, ...interface{}) RowAbstract
		ExecContext(context.Context, string, ...interface{}) (ResultAbstract, error)
		BeginTx(context.Context, *sql.TxOptions) (TxAbstract, error)
	}

	TxAbstract interface {
		Commit() error
		QueryContext(context.Context, string, ...interface{}) (RowsAbstract, error)
		QueryRowContext(context.Context, string, ...interface{}) RowAbstract
		ExecContext(context.Context, string, ...interface{}) (ResultAbstract, error)
		Rollback() error
	}

	RowAbstract interface {
		Err() error
		Scan(...interface{}) error
	}

	RowsAbstract interface {
		Err() error
		Scan(...interface{}) error
		Next() bool
		Close() error
	}

	ResultAbstract interface {
		// LastInserted() (int, error)
		RowsAffected() (int, error)
	}
)
