package migration

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
	"github.com/maru44/enva/service/api/pkg/config"
	"github.com/maru44/perr"

	// required for using migrate package.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Psql struct {
	conn *sql.DB
	url  string
}

const (
	maxLifeTime = 5 * time.Minute
)

func NewDB(ctx context.Context) (*Psql, error) {
	d := &Psql{
		url: config.POSTGRES_URL,
	}
	if err := d.connect(); err != nil {
		return nil, perr.Wrap(err, perr.InternalServerError)
	}
	return d, nil
}

func (d *Psql) connect() error {
	conn, err := sql.Open("postgres", d.url)
	if err != nil {
		return perr.Wrap(err, perr.InternalServerError)
	}

	conn.SetMaxOpenConns(config.POSTGRES_MAX_CONNECTIONS)
	conn.SetMaxIdleConns(config.POSTGRES_MAX_IDLE_CONNECTIONS)
	conn.SetConnMaxLifetime(maxLifeTime)

	d.conn = conn
	return nil
}

//go:embed postgres/*.sql
var migrationFileFS embed.FS

func (d *Psql) Up(ctx context.Context) error {
	return d.useMigrator(ctx, func(m *migrate.Migrate) error {
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		return nil
	})
}

func (d *Psql) Down(ctx context.Context) error {
	return d.useMigrator(ctx, func(m *migrate.Migrate) error {
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		return nil
	})
}

func (d *Psql) Drop(ctx context.Context) error {
	return d.useMigrator(ctx, func(m *migrate.Migrate) error {
		if err := m.Drop(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		return nil
	})
}

func (d *Psql) Version(ctx context.Context) (*uint, bool, error) {
	var (
		version *uint
		dirty   bool
	)
	if err := d.useMigrator(ctx, func(m *migrate.Migrate) error {
		v, d, err := m.Version()
		if err != nil {
			return err
		}
		version = &v
		dirty = d
		return nil
	}); err != nil {
		return nil, false, err
	}
	return version, dirty, nil
}

func (d *Psql) VersionDown(ctx context.Context) error {
	return d.useMigrator(ctx, func(m *migrate.Migrate) error {
		v, d, err := m.Version()
		if err != nil {
			return err
		}
		if d {
			if err := m.Force(int(v - 1)); err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *Psql) useMigrator(ctx context.Context, f func(*migrate.Migrate) error) (err error) {
	drv, err := iofs.New(migrationFileFS, "postgres")
	if err != nil {
		return perr.Wrap(err, perr.InternalServerError)
	}
	m, err := migrate.NewWithSourceInstance("iofs", drv, d.url)
	if err != nil {
		return perr.Wrap(err, perr.InternalServerError)
	}
	m.Log = newLogger(ctx)

	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			err = perr.Wrap(srcErr, perr.InternalServerError)
		}
		if dbErr != nil {
			err = perr.Wrap(dbErr, perr.InternalServerError)
		}
		connectErr := d.connect()
		if connectErr != nil {
			err = perr.Wrap(connectErr, perr.InternalServerError)
		}
	}()

	err = f(m)
	return err
}
