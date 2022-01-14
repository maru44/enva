package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/maru44/enva/service/api/internal/config"
	"github.com/maru44/enva/service/api/internal/interface/database"
)

type (
	SqlHandler struct {
		Conn *sql.DB
	}

	SqlRow struct {
		Row *sql.Row
	}

	SqlRows struct {
		Rows *sql.Rows
	}

	SqlResult struct {
		Result sql.Result
	}

	SqlTx struct {
		Tx *sql.Tx
	}
)

const (
	maxLifeTime = 5 * time.Minute
)

/***************************
        instance
***************************/

func NewSqlHandler() database.ISqlHandler {
	conn, err := sql.Open("postgres", fmt.Sprintf(config.POSTGRES_URL))
	if err != nil {
		panic(err)
	}
	conn.SetMaxOpenConns(config.POSTGRES_MAX_CONNECTIONS)
	conn.SetMaxIdleConns(config.POSTGRES_MAX_IDLE_CONNECTIONS)
	conn.SetConnMaxLifetime(maxLifeTime)

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

/***************************
    handler's methods
***************************/

func (h *SqlHandler) BeginTx(ctx context.Context, opts *sql.TxOptions) (database.ITx, error) {
	t, err := h.Conn.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	tx := &SqlTx{
		Tx: t,
	}
	return tx, err
}

func (h *SqlHandler) QueryContext(ctx context.Context, st string, args ...interface{}) (database.IRows, error) {
	rows, err := h.Conn.QueryContext(ctx, st, args...)
	if err != nil {
		return nil, err
	}

	r := &SqlRows{
		Rows: rows,
	}
	return r, nil
}

func (h *SqlHandler) QueryRowContext(ctx context.Context, st string, args ...interface{}) database.IRow {
	row := h.Conn.QueryRowContext(ctx, st, args...)
	r := &SqlRow{
		Row: row,
	}
	return r
}

func (h *SqlHandler) ExecContext(ctx context.Context, st string, args ...interface{}) (database.IResult, error) {
	result, err := h.Conn.ExecContext(ctx, st, args...)
	if err != nil {
		return nil, err
	}

	r := &SqlResult{
		Result: result,
	}
	return r, nil
}

/***************************
    transaction's methods
***************************/

func (t SqlTx) Commit() error {
	return t.Tx.Commit()
}

func (t SqlTx) Rollback() {
	if err := t.Tx.Rollback(); err != nil {
		log.Fatalf("failed to rollback: %v", err)
	}
	return
}

func (t SqlTx) QueryContext(ctx context.Context, st string, args ...interface{}) (database.IRows, error) {
	rows, err := t.Tx.QueryContext(ctx, st, args...)
	if err != nil {
		return nil, err
	}

	r := &SqlRows{
		Rows: rows,
	}
	return r, nil
}

func (t SqlTx) QueryRowContext(ctx context.Context, st string, args ...interface{}) database.IRow {
	row := t.Tx.QueryRowContext(ctx, st, args...)
	r := &SqlRow{
		Row: row,
	}
	return r
}

func (t SqlTx) ExecContext(ctx context.Context, st string, args ...interface{}) (database.IResult, error) {
	stmt, err := t.Tx.Prepare(st)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	exe, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	r := &SqlResult{
		Result: exe,
	}
	return r, nil
}

/***************************
    Rows' methods
***************************/

func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRows) Err() error {
	return r.Rows.Err()
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}

/***************************
    Row's methods
***************************/

func (r SqlRow) Err() error {
	return r.Row.Err()
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

/***************************
    Result's methods
***************************/

func (r SqlResult) RowsAffected() (int, error) {
	affected, err := r.Result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}
