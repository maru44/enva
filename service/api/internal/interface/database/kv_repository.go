package database

import (
	"context"
	"fmt"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/pkg/domain"
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

func (repo *KvRepository) DetailValid(ctx context.Context, key domain.KvKey, projectID domain.ProjectID) (*domain.Kv, error) {
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

func (repo *KvRepository) Create(ctx context.Context, input domain.KvInput, projectID domain.ProjectID) (*domain.KvID, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	if err := input.Validate(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	// if key exists >> return error
	var preId string
	row := repo.QueryRowContext(ctx, queryset.ValidKvDetailID, input.Key, projectID)
	if err := row.Scan(&preId); err == nil {
		return nil, perr.New(fmt.Sprintf("the key is already exists: %s", input.Key), perr.BadRequest)
	}

	var id string
	if err := repo.QueryRowContext(
		ctx,
		queryset.KvInsertQuery,
		input.Key, input.Value, projectID, user.ID,
	).Scan(&id); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	ID := domain.KvID(id)
	return &ID, nil
}

func (repo *KvRepository) Update(ctx context.Context, input domain.KvInput, projectID domain.ProjectID) (*domain.KvID, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	if err := input.Validate(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	tx, err := repo.BeginTx(ctx, nil)
	if err != nil {
		return nil, perr.Wrap(err, perr.InternalServerErrorWithUrgency)
	}

	// deactivate existing kv
	exe, err := tx.ExecContext(
		ctx,
		queryset.KvDeactivateQuery,
		user.ID, projectID, input.Key,
	)
	if err != nil {
		tx.Rollback()
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	if affected, err := exe.RowsAffected(); err != nil {
		tx.Rollback()
		return nil, perr.Wrap(err, perr.BadRequest)
	} else if err == nil && affected == 0 {
		// @INFO if want to upsert remove this condition
		return nil, perr.New("the key does not exists in this project", perr.NotFound, "the key does not exists in this project")
	}

	// create new kv
	var id string
	if err := tx.QueryRowContext(
		ctx,
		queryset.KvInsertQuery,
		input.Key, input.Value, projectID, user.ID,
	).Scan(&id); err != nil {
		tx.Rollback()
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	tx.Commit()

	ID := domain.KvID(id)
	return &ID, nil
}

func (repo *KvRepository) Delete(ctx context.Context, kvID domain.KvID, projectID domain.ProjectID) (int, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	// deactivate existing kv
	exe, err := repo.ExecContext(
		ctx,
		queryset.KvDeactivateByIdQuery,
		user.ID, projectID, kvID,
	)
	if err != nil {
		return 0, perr.Wrap(err, perr.BadRequest)
	}

	affected, err := exe.RowsAffected()
	if err != nil {
		return 0, perr.Wrap(err, perr.BadRequest)
	} else if affected == 0 {
		return 0, perr.New("No result", perr.BadRequest)
	}
	return affected, nil
}

func (repo *KvRepository) DeleteByKey(ctx context.Context, key domain.KvKey, projectID domain.ProjectID) (int, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	// deactivate existing kv
	exe, err := repo.ExecContext(
		ctx,
		queryset.KvDeactivateQuery,
		user.ID, projectID, key,
	)
	if err != nil {
		return 0, perr.Wrap(err, perr.BadRequest)
	}

	affected, err := exe.RowsAffected()
	if err != nil {
		return 0, perr.Wrap(err, perr.BadRequest)
	} else if affected == 0 {
		return 0, perr.New("No result", perr.BadRequest)
	}
	return affected, nil
}
