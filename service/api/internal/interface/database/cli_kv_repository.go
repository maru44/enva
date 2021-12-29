package database

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type CliKvRepository struct {
	ISqlHandler
}

func (repo *CliKvRepository) BulkInsert(ctx context.Context, projectID domain.ProjectID, inputs []domain.KvInput) error {
	user, err := domain.UserFromCtx(ctx)

	tx, err := repo.BeginTx(ctx, nil)
	if err != nil {
		return perr.Wrap(err, perr.InternalServerError)
	}

	for _, in := range inputs {
		if err := in.Validate(); err != nil {
			tx.Rollback()
			return perr.Wrap(err, perr.BadRequest)
		}

		var id string
		if err := tx.QueryRowContext(
			ctx,
			queryset.KvInsertQuery,
			in.Key, in.Value, projectID, user.ID,
		).Scan(&id); err != nil {
			tx.Rollback()
			return perr.Wrap(err, perr.BadRequest)
		}
	}
	tx.Commit()

	return nil
}
