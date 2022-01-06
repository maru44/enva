package database

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgRepository struct {
	ISqlHandler
}

func (repo *OrgRepository) List(ctx context.Context) ([]domain.Org, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.OrgListQuery, user.ID,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var orgs []domain.Org
	for rows.Next() {
		var (
			o        domain.Org
			userType domain.UserType
		)
		if err := rows.Scan(
			&o.ID, &o.Slug, &o.Name, &o.Description,
			&o.CreatedAt, &o.UpdatedAt, &userType,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}

		orgs = append(orgs, o)
	}
	return orgs, nil
}

func (repo *OrgRepository) ListOwnerAdmin(ctx context.Context) ([]domain.Org, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.OrgListQuery, user.ID,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var orgs []domain.Org
	for rows.Next() {
		var (
			o        domain.Org
			userType domain.UserType
		)
		if err := rows.Scan(
			&o.ID, &o.Slug, &o.Name, &o.Description,
			&o.CreatedAt, &o.UpdatedAt, &userType,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}

		if userType == domain.UserTypeOwner || userType == domain.UserTypeAdmin {
			orgs = append(orgs, o)
		}
	}
	return orgs, nil
}

func (repo *OrgRepository) Detail(ctx context.Context, orgID domain.OrgID) (*domain.Org, *domain.UserType, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}
	row := repo.QueryRowContext(
		ctx,
		queryset.OrgDetailQuery,
		user.ID, orgID,
	)
	if err := row.Err(); err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}
	var (
		o       *domain.Org
		ownerID domain.UserID
		ut      *domain.UserType
	)
	if err := row.Scan(
		&o.ID, &o.Slug, &o.Name, &o.Description,
		&ownerID, &o.CreatedAt, &o.UpdatedAt, &o.UserCount,
		&ut,
	); err != nil {
		return nil, nil, perr.Wrap(err, perr.Wrap(err, perr.NotFound))
	}

	o.CreatedBy = domain.User{ID: ownerID}
	return o, ut, nil
}

func (repo *OrgRepository) DetailBySlug(ctx context.Context, slug string) (*domain.Org, *domain.UserType, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}
	row := repo.QueryRowContext(
		ctx,
		queryset.OrgDetailBySlugQuery,
		user.ID, slug,
	)
	if err := row.Err(); err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}
	var (
		o  *domain.Org = &domain.Org{}
		u  domain.User
		ut *domain.UserType
	)
	if err := row.Scan(
		&o.ID, &o.Slug, &o.Name, &o.Description, &o.IsValid,
		&u.ID, &o.CreatedAt, &o.UpdatedAt, &o.UserCount,
		&ut,
	); err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}

	o.CreatedBy = u
	return o, ut, nil
}

func (repo *OrgRepository) Create(ctx context.Context, input domain.OrgInput) (*string, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	// @TODO confirm org count of owner

	tx, err := repo.BeginTx(ctx, nil)
	if err != nil {
		return nil, perr.Wrap(err, perr.InternalServerError)
	}

	var id, slug *string
	if err := tx.QueryRowContext(ctx,
		queryset.OrgCreateQuery,
		input.Slug, input.Name, input.Description, user.ID,
	).Scan(&id, &slug); err != nil {
		tx.Rollback()
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	var memberID *string
	if err := tx.QueryRowContext(ctx,
		queryset.RelOrgMembersInsertQuery,
		id, user.ID, domain.UserTypeOwner, nil,
	).Scan(&memberID); err != nil {
		tx.Rollback()
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	tx.Commit()

	return slug, nil
}
