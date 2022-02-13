package database

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database/qs"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type CliKvRepository struct {
	ISqlHandler
}

func (repo *CliKvRepository) BulkInsert(ctx context.Context, projectID domain.ProjectID, inputs []domain.KvInput) error {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.ErrForbidden)
	}

	tx, err := repo.BeginTx(ctx, nil)
	if err != nil {
		return perr.Wrap(err, perr.ErrInternalServerError)
	}

	for _, in := range inputs {
		if err := in.Validate(); err != nil {
			tx.Rollback()
			return perr.Wrap(err, perr.ErrBadRequest)
		}

		var id string
		if err := tx.QueryRowContext(
			ctx,
			qs.KvInsertQuery,
			in.Key, in.Value, projectID, user.ID,
		).Scan(&id); err != nil {
			tx.Rollback()
			return perr.Wrap(err, perr.ErrBadRequest)
		}
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return perr.Wrap(err, perr.ErrInternalServerError)
	}

	return nil
}
