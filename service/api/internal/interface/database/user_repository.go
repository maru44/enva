package database

import (
	"context"
	"database/sql"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/internal/interface/password"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/tools"
	"github.com/maru44/perr"
)

type UserRepository struct {
	ISqlHandler
	password.IPassword
}

func (repo *UserRepository) GetByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	row := repo.QueryRowContext(ctx, queryset.UserGetByIDQuery, id)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var u *domain.User
	if err := row.Scan(
		&u.ID, &u.Email, &u.Username, &u.ImageURL, &u.CliPassword,
		&u.IsValid, &u.IsEmailVerified, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	if u.CliPassword != nil {
		u.HasCliPassword = true
	}

	rows, err := repo.QueryContext(ctx, queryset.UsersBelongOrgQuery, id)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	for rows.Next() {
		var (
			o        domain.Org
			userType domain.UserType
		)
		if err := rows.Scan(
			&o.ID, &o.Slug, &o.Name, &userType,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}

		if userType == domain.UserTypeOwner {
			u.OwnerOf = append(u.OwnerOf, o)
		}
		if userType == domain.UserTypeAdmin {
			u.AdminOf = append(u.AdminOf, o)
		}
		if userType == domain.UserTypeUser {
			u.UserOf = append(u.UserOf, o)
		}
	}

	return u, nil
}

func (repo *UserRepository) Create(ctx context.Context) (*string, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}
	input := &domain.UserInput{
		ID:              user.ID.String(),
		Email:           user.Email,
		Username:        user.Username,
		IsEmailVerified: user.IsEmailVerified,
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	var id *string
	row := repo.QueryRowContext(ctx, queryset.UserExistsQuery, input.ID)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.InternalServerError)
	}
	err = row.Scan(&id)
	if err == nil {
		return nil, perr.New("The user already exists", perr.BadRequest, "The user already exists")
	}
	if err != sql.ErrNoRows {
		return nil, perr.Wrap(err, perr.InternalServerError)
	}

	if err := repo.QueryRowContext(
		ctx,
		queryset.UserInsertQuery,
		input.ID, input.Email, input.Username, input.IsEmailVerified,
	).Scan(&id); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return id, nil
}

func (repo *UserRepository) UpdateCliPassword(ctx context.Context) (*string, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	// generate new password and its hash
	rawPass := tools.GenRandSlug(48)
	hashed, err := repo.Generate(rawPass)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	input := &domain.UserCliPasswordInput{
		ID:          user.ID.String(),
		CliPassword: hashed,
	}
	if err := input.Validate(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	exe, err := repo.ExecContext(ctx,
		queryset.UserUpdateCliPasswordQuery,
		input.CliPassword, input.ID,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	if affected, err := exe.RowsAffected(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	} else if affected == 0 {
		return nil, perr.New("There is no affected row", perr.BadRequest)
	}

	return &rawPass, nil
}

func (repo *UserRepository) GetUserCli(ctx context.Context, input *domain.UserCliValidationInput) (*domain.User, error) {
	row := repo.QueryRowContext(ctx,
		queryset.UserGetByEmailOrPassword,
		input.EmailOrUsername,
	)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	u := &domain.User{}
	if err := row.Scan(
		&u.ID, &u.Email, &u.Username, &u.ImageURL,
		&u.CliPassword, &u.CreatedAt, &u.UpdatedAt,
	); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	if err := repo.Check(*u.CliPassword, input.CliPassword); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return u, nil
}
