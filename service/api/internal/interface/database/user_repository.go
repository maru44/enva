package database

import (
	"context"
	"database/sql"

	"github.com/maru44/enva/service/api/internal/interface/database/qs"
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
	row := repo.QueryRowContext(ctx, qs.UserGetByIDQuery, id)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	u := &domain.User{}
	var sID, sSubScriptionID, sCustomerID, sProductID, sSubscriptionStatus *string
	if err := row.Scan(
		&u.ID, &u.Email, &u.Username, &u.ImageURL, &u.CliPassword,
		&u.IsValid, &u.IsEmailVerified, &u.CreatedAt, &u.UpdatedAt,
		&sID, &sSubScriptionID, &sCustomerID,
		&sProductID, &sSubscriptionStatus,
	); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	if u.CliPassword != nil {
		u.HasCliPassword = true
	}

	if sID != nil {
		sub := &domain.Subscription{
			ID:                       *sID,
			StripeSubscriptionID:     *sSubScriptionID,
			StripeCustomerID:         *sCustomerID,
			StripeProductID:          *sProductID,
			StripeSubscriptionStatus: *sSubscriptionStatus,
		}
		u.Subscription = sub
	}

	return u, nil
}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := repo.QueryRowContext(ctx, qs.UserGetByEmailQuery, email)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	u := &domain.User{}
	var sID, sSubScriptionID, sCustomerID, sProductID, sSubscriptionStatus *string
	if err := row.Scan(
		&u.ID, &u.Email, &u.Username, &u.ImageURL, &u.CliPassword,
		&u.IsValid, &u.IsEmailVerified, &u.CreatedAt, &u.UpdatedAt,
		&sID, &sSubScriptionID, &sCustomerID,
		&sProductID, &sSubscriptionStatus,
	); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	if u.CliPassword != nil {
		u.HasCliPassword = true
	}

	if sID != nil {
		sub := &domain.Subscription{
			ID:                       *sID,
			StripeSubscriptionID:     *sSubScriptionID,
			StripeCustomerID:         *sCustomerID,
			StripeProductID:          *sProductID,
			StripeSubscriptionStatus: *sSubscriptionStatus,
		}
		u.Subscription = sub
	}

	return u, nil
}

func (repo *UserRepository) UpsertIfNotInvalid(ctx context.Context) (*string, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}
	input := &domain.UserInput{
		ID:              user.ID.String(),
		Email:           user.Email,
		Username:        user.Username,
		IsEmailVerified: user.IsEmailVerified,
		ImageURL:        user.ImageURL,
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	var (
		id, imageURL             *string
		isValid, isEmailVerified bool
	)
	row := repo.QueryRowContext(ctx, qs.UserExistsQuery, input.ID)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.InternalServerError)
	}

	err = row.Scan(&id, &isValid, &isEmailVerified, &imageURL)
	// if ex
	if err == nil {
		// if not valid return 403
		if !isValid {
			return nil, perr.New("user is not valid", perr.Forbidden, "user is not valid")
		}

		// update isEmailVerified and imageURL if it changed
		isNeedToUpdate := false
		switch {
		case input.ImageURL == nil && imageURL == nil:
		case input.IsEmailVerified != isEmailVerified,
			input.ImageURL == nil && imageURL != nil,
			input.ImageURL != nil && imageURL == nil,
			*input.ImageURL != *imageURL:
			isNeedToUpdate = true
		}
		if isNeedToUpdate {
			exe, err := repo.ExecContext(ctx,
				qs.UserUpdateImageOrIsEmailVerifiedQuery,
				input.IsEmailVerified, input.ImageURL, input.ID,
			)
			if err != nil {
				return nil, perr.Wrap(err, perr.BadRequest)
			}

			affected, err := exe.RowsAffected()
			if err != nil {
				return nil, perr.Wrap(err, perr.BadRequest)
			}
			if affected == 0 {
				return nil, perr.New("There is no affected row", perr.BadRequest)
			}
		}
		return id, nil
	}
	if err != sql.ErrNoRows {
		return nil, perr.Wrap(err, perr.InternalServerError)
	}

	// if not ex
	if err := repo.QueryRowContext(
		ctx,
		qs.UserInsertQuery,
		input.ID, input.Email, input.Username, input.IsEmailVerified, input.ImageURL,
	).Scan(&id); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return id, nil
}

func (repo *UserRepository) UpdateValid(ctx context.Context, input domain.UserUpdateIsValidInput) error {
	if err := input.Validate(); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	exe, err := repo.ExecContext(ctx,
		qs.UserUpdateValidQuery,
		input.IsValid, input.ID,
	)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	if affected, err := exe.RowsAffected(); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	} else if affected == 0 {
		return perr.New("There is no affected row", perr.BadRequest)
	}
	return nil
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
		ID:          user.ID,
		CliPassword: hashed,
	}
	if err := input.Validate(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	exe, err := repo.ExecContext(ctx,
		qs.UserUpdateCliPasswordQuery,
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
		qs.UserGetByEmailAdnPassword,
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
