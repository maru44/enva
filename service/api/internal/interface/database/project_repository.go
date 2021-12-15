package database

import (
	"context"

	"github.com/maru44/ichigo/service/api/internal/interface/database/queryset"
	"github.com/maru44/ichigo/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type ProjectReposotory struct {
	ISqlHandler
}

func (repo *ProjectReposotory) ListByUser(ctx context.Context) ([]domain.Project, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	rows, err := repo.QueryContext(ctx, queryset.ProjectListByUserQuery, user.ID)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var ps []domain.Project
	for rows.Next() {
		var (
			p             domain.Project
			userID, orgID *string
		)
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
			&userID, &orgID,
			&p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}

		// set user
		if userID != nil {
			p.OwnerUser = &user
		}

		// set org
		if orgID != nil {
			p.OwnerOrg = &domain.Org{
				ID: domain.OrgID(*orgID),
			}
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (repo *ProjectReposotory) ListByOrg(ctx context.Context, orgID domain.OrgID) ([]domain.Project, error) {
	rows, err := repo.QueryContext(ctx, queryset.ProjectListByOrgQuery, orgID)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var ps []domain.Project
	for rows.Next() {
		var (
			p     domain.Project
			orgID *string
		)
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
			&orgID,
			&p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}

		// set org
		if orgID != nil {
			p.OwnerOrg = &domain.Org{
				ID: domain.OrgID(*orgID),
			}
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (repo *ProjectReposotory) SlugListByUser(ctx context.Context) ([]string, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	rows, err := repo.QueryContext(ctx, queryset.ProjectSlugListByUserQuery, user.ID)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var slugs []string
	for rows.Next() {
		var (
			slug string
		)
		if err := rows.Scan(
			&slug,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}

		slugs = append(slugs, slug)
	}

	return slugs, nil
}

func (repo *ProjectReposotory) GetBySlug(ctx context.Context, slug string) (*domain.Project, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	row := repo.QueryRowContext(ctx, queryset.ProjectDetailBySlugQuery, slug)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var (
		p             *domain.Project = &domain.Project{}
		userID, orgID *string
	)
	if err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
		&userID, &orgID,
		&p.IsValid, &p.IsDeleted,
		&p.CreatedAt, &p.UpdatedAt,
	); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	// set user
	if userID != nil {
		if *userID != user.ID {
			return nil, perr.New(perr.NotFound.Error(), perr.NotFound)
		}
		p.OwnerUser = &user
	}

	// set org
	if orgID != nil {
		// @TODO validate user is
		p.OwnerOrg = &domain.Org{
			ID: domain.OrgID(*orgID),
		}
	}

	return p, nil
}

func (repo *ProjectReposotory) GetByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)

	row := repo.QueryRowContext(ctx, queryset.ProjectDetailByIDQuery, id)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var (
		p             *domain.Project = &domain.Project{}
		userID, orgID *string
	)
	if err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
		&userID, &orgID,
		&p.IsValid, &p.IsDeleted,
		&p.CreatedAt, &p.UpdatedAt,
	); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	// set user
	if userID != nil {
		if *userID != user.ID {
			return nil, perr.New(perr.NotFound.Error(), perr.NotFound)
		}
		p.OwnerUser = &user
	}

	// set org
	if orgID != nil {
		// @TODO validate user is
		p.OwnerOrg = &domain.Org{
			ID: domain.OrgID(*orgID),
		}
	}

	return p, nil
}

func (repo *ProjectReposotory) Create(ctx context.Context, input domain.ProjectInput) (*string, error) {
	user := ctx.Value(domain.CtxUserKey).(domain.User)
	var (
		in           domain.ProjectInput
		inputU, slug *string
	)

	ownerType := domain.OwnerTypeUser
	if in.OrgID != nil {
		ownerType = domain.OwnerTypeOrg
	} else {
		inputU = &user.ID
	}

	// @TODO get user's or org's project slug(all)

	if err := repo.QueryRowContext(
		ctx,
		queryset.ProjectCreateQuery,
		input.Name, input.Slug, input.Description, ownerType, inputU, in.OrgID,
	).Scan(&slug); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return slug, nil
}
