package database

import (
	"context"

	"github.com/maru44/ichigo/service/api/internal/interface/database/queryset"
	"github.com/maru44/ichigo/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type KvRepository struct {
	ISqlHandler
}

func (repo *KvRepository) ListValid(ctx context.Context, projectID domain.ProjectID) (kvs []domain.Kv, err error) {
	rows, err := repo.QueryContext(ctx, queryset.ValidKvListOfProject, projectID)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if rows.Err() != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	for rows.Next() {
		k := domain.Kv{}
		err = rows.Scan(
			&k.ID, &k.Key,
			&k.Value, &k.IsValid,
			&k.CreatedAt, &k.UpdatedAt,
		)
		if err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}
		k.ProjectID = projectID
		kvs = append(kvs, k)
	}

	return kvs, nil
}

// DetailValid(context.Context, domain.KvKey, domain.ProjectID) (*domain.Kv, error)
func (repo *KvRepository) DetailValid(ctx context.Context, key domain.KvKey, projectID domain.ProjectID) (*domain.Kv, error) {
	// user := ctx.Value(domain.CtxUserKey).(domain.User)

	row := repo.QueryRowContext(ctx, queryset.ValidKvDetail, key, projectID)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	kv := &domain.Kv{}
	if err := row.Scan(
		&kv.ID, &kv.Key, &kv.Value, &kv.IsValid,
		&kv.CreatedAt, &kv.UpdatedAt,
	); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	return kv, nil
}

func (repo *KvRepository) Create(ctx context.Context, input domain.KvInput, projectID domain.ProjectID) (id string, err error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	tx, err := repo.BeginTx(ctx, nil)
	if err != nil {
		return "", perr.Wrap(err, perr.InternalServerErrorWithUrgency)
	}

	row := tx.QueryRowContext(ctx, queryset.ValidKvDetailID, input.Key, projectID)
	// if already exists key in projects
	// update its is_valid and updated_by
	if err := row.Err(); err == nil {
		var kvID string
		if err := row.Scan(&kvID); err != nil {
			tx.Rollback()
			return id, perr.Wrap(err, perr.BadRequest)
		}

		exe, err := repo.ExecContext(
			ctx,
			queryset.KvDeactivateQuery,
			user, kvID,
		)
		if err != nil {
			tx.Rollback()
			return id, perr.Wrap(err, perr.BadRequest)
		}

		if _, err := exe.RowsAffected(); err != nil {
			tx.Rollback()
			return id, perr.Wrap(err, perr.BadRequest)
		}
	}

	if err := repo.QueryRowContext(
		ctx,
		queryset.KvInsertQuery,
		input.Key, input.Value, projectID, user.ID,
	).Scan(&id); err != nil {
		tx.Rollback()
		return id, perr.Wrap(err, perr.BadRequest)
	}

	tx.Commit()

	return id, nil
}
