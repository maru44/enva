package database

import (
	"context"
	"database/sql"
)

type (
	ISqlHandler interface {
		QueryContext(context.Context, string, ...interface{}) (IRows, error)
		QueryRowContext(context.Context, string, ...interface{}) IRow
		ExecContext(context.Context, string, ...interface{}) (IResult, error)
		BeginTx(context.Context, *sql.TxOptions) (ITx, error)
	}

	ITx interface {
		Commit() error
		QueryContext(context.Context, string, ...interface{}) (IRows, error)
		QueryRowContext(context.Context, string, ...interface{}) IRow
		ExecContext(context.Context, string, ...interface{}) (IResult, error)
		Rollback() error
	}

	IRow interface {
		Err() error
		Scan(...interface{}) error
	}

	IRows interface {
		Err() error
		Scan(...interface{}) error
		Next() bool
		Close() error
	}

	IResult interface {
		// LastInserted() (int, error)
		RowsAffected() (int, error)
	}
)
