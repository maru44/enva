package database

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/internal/interface/password"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/tools"
	"github.com/maru44/perr"
)

type CliUserRepository struct {
	ISqlHandler
	password.IPassword
}

func (repo *CliUserRepository) Create(ctx context.Context) (string, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)
	var id string

	// gen password
	rawPass := tools.GenRandSlug(48)
	hashed, err := repo.Generate(rawPass)
	if err != nil {
		return "", perr.Wrap(err, perr.BadRequest)
	}

	input, err := domain.CreateAndValidateCliUserCraeteInput(user, hashed)
	if err != nil {
		return "", perr.Wrap(err, perr.BadRequest)
	}

	if err := repo.QueryRowContext(
		ctx,
		queryset.CliUserInsertQuery,
		input.Email, input.Username, input.Password,
	).Scan(&id); err != nil {
		return "", perr.Wrap(err, perr.BadRequest)
	}

	return rawPass, nil
}

func (repo *CliUserRepository) Update(ctx context.Context) (string, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	// gen password
	rawPass := tools.GenRandSlug(48)
	hashed, err := repo.Generate(rawPass)
	if err != nil {
		return "", perr.Wrap(err, perr.BadRequest)
	}

	input, err := domain.CreateAndValidateCliUserCraeteInput(user, hashed)
	if err != nil {
		return "", perr.Wrap(err, perr.BadRequest)
	}

	exe, err := repo.ExecContext(ctx,
		queryset.CliUserUpdateQuery,
		input.Password, input.Email,
	)
	if err != nil {
		return "", perr.Wrap(err, perr.BadRequest)
	}

	if affected, err := exe.RowsAffected(); err != nil {
		return "", perr.Wrap(err, perr.BadRequest)
	} else if affected == 0 {
		return "", perr.New("there is no affected row", perr.BadRequest)
	}

	return rawPass, nil
}

func (repo *CliUserRepository) Validate(ctx context.Context, input *domain.CliUserValidateInput) error {
	row := repo.QueryRowContext(ctx,
		queryset.CliUserGetPasswordByEmailOrPassword,
		input.EmailOrUsername,
	)
	if err := row.Err(); err != nil {
		return perr.Wrap(err, perr.NotFound)
	}

	var dbPass string
	if err := row.Scan(&dbPass); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	if err := repo.Check(dbPass, input.Password); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}

func (repo *CliUserRepository) Exists(ctx context.Context) error {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	row := repo.QueryRowContext(ctx,
		queryset.CliUserExistsQuery,
		user.Email,
	)
	if err := row.Err(); err != nil {
		return perr.Wrap(err, perr.NotFound)
	}

	var id string
	if err := row.Scan(&id); err != nil {
		return perr.Wrap(err, perr.NotFound)
	}

	return nil
}

// @TODO implement
func (repo *CliUserRepository) Delete(ctx context.Context) error {
	return nil
}

func (repo *CliUserRepository) GetUser(ctx context.Context, input *domain.CliUserValidateInput) (*domain.User, error) {
	row := repo.QueryRowContext(ctx,
		queryset.CliUserGet,
		input.EmailOrUsername,
	)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var (
		dbPass string
		c      = domain.CliUser{}
	)
	if err := row.Scan(
		&c.CognitoID, &c.Email, &c.Username,
		&dbPass,
	); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	if err := repo.Check(dbPass, input.Password); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return c.ToUser(), nil
}
