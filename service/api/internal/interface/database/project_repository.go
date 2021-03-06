package database

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"github.com/maru44/enva/service/api/internal/interface/database/qs"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/enva/service/api/pkg/tools"
	"github.com/maru44/perr"
)

type ProjectReposotory struct {
	ISqlHandler
}

func (repo *ProjectReposotory) ListAll(ctx context.Context) ([]domain.Project, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	orgRows, err := repo.QueryContext(ctx,
		qs.OrgValidListQuery,
		user.ID,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrBadRequest)
	}
	if err := orgRows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrBadRequest)
	}

	var orgs []domain.Org
	for orgRows.Next() {
		var (
			o        domain.Org
			userType domain.UserType
		)
		if err := orgRows.Scan(
			&o.ID, &o.Slug, &o.Name, &o.Description,
			&o.CreatedAt, &o.UpdatedAt, &userType,
		); err != nil {
			return nil, perr.Wrap(err, perr.ErrBadRequest)
		}
		orgs = append(orgs, o)
	}

	var projects []domain.Project
	for i, o := range orgs {
		ps, err := repo.ListByOrg(ctx, o.ID)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrBadRequest)
		}
		// set org
		for _, p := range ps {
			// &o is failed @TODO research
			p.OwnerOrg = &orgs[i]
			projects = append(projects, p)
		}
	}

	ps, err := repo.ListByUser(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrBadRequest)
	}
	projects = append(projects, ps...)

	sort.Slice(projects, func(i, j int) bool { return projects[i].UpdatedAt.String() < projects[j].UpdatedAt.String() })

	return projects, nil
}

func (repo *ProjectReposotory) ListByUser(ctx context.Context) ([]domain.Project, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	rows, err := repo.QueryContext(ctx, qs.ProjectValidListByUserQuery, user.ID)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
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
			return nil, perr.Wrap(err, perr.ErrNotFound)
		}

		// set user
		if userID != nil {
			p.OwnerUser = user
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
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	rows, err := repo.QueryContext(ctx,
		qs.ProjectListValidByOrgQuery,
		orgID, user.ID,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	var ps []domain.Project
	for rows.Next() {
		var (
			p      domain.Project
			userID *string
			orgID  *string
		)
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
			&userID, &orgID,
			&p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, perr.Wrap(err, perr.ErrNotFound)
		}

		// set org
		if orgID != nil {
			p.OwnerOrg = &domain.Org{
				ID: domain.OrgID(*orgID),
			}
		}

		if userID != nil {
			p.OwnerUser = &domain.User{
				ID: domain.UserID(*userID),
			}
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (repo *ProjectReposotory) SlugListByUser(ctx context.Context) ([]string, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	rows, err := repo.QueryContext(ctx, qs.ProjectValidSlugListByUserQuery, user.ID)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	var slugs []string
	for rows.Next() {
		var (
			slug string
		)
		if err := rows.Scan(
			&slug,
		); err != nil {
			return nil, perr.Wrap(err, perr.ErrNotFound)
		}

		slugs = append(slugs, slug)
	}

	return slugs, nil
}

func (repo *ProjectReposotory) GetBySlug(ctx context.Context, slug string) (*domain.Project, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	row := repo.QueryRowContext(ctx, qs.ProjectValidDetailBySlugQuery, slug, user.ID)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	var (
		p      *domain.Project = &domain.Project{}
		userID domain.UserID
		orgID  *string
	)
	if err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
		&userID, &orgID,
		&p.IsValid, &p.DeletedAt,
		&p.CreatedAt, &p.UpdatedAt,
	); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	// set user
	if userID != "" {
		if userID != user.ID {
			return nil, perr.New(perr.ErrNotFound.Error(), perr.ErrNotFound)
		}
		p.OwnerUser = user
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

func (repo *ProjectReposotory) GetBySlugAndOrgID(ctx context.Context, slug string, orgID domain.OrgID) (*domain.Project, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	row := repo.QueryRowContext(ctx, qs.ProjectValidDetailBySlugAndOrgIdQuery, user.ID, orgID, slug)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	var (
		p      *domain.Project = &domain.Project{}
		userID *domain.UserID
		o      *domain.Org = &domain.Org{ID: orgID}
	)
	if err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
		&userID,
		&p.IsValid, &p.DeletedAt,
		&p.CreatedAt, &p.UpdatedAt,
		&o.Slug, &o.Name,
	); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	// set user
	if userID != nil {
		p.OwnerUser = &domain.User{ID: *userID}
	}

	p.OwnerOrg = o

	return p, nil
}

func (repo *ProjectReposotory) GetBySlugAndOrgSlug(ctx context.Context, slug, orgSlug string) (*domain.Project, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	row := repo.QueryRowContext(ctx, qs.ProjectValidDetailBySlugAndOrgSlugQuery, user.ID, orgSlug, slug)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	var (
		p      *domain.Project = &domain.Project{}
		userID *domain.UserID
		o      *domain.Org = &domain.Org{Slug: orgSlug}
	)
	if err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
		&userID,
		&p.IsValid, &p.DeletedAt,
		&p.CreatedAt, &p.UpdatedAt,
		&o.ID, &o.Name,
	); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	// set user
	if userID != nil {
		p.OwnerUser = &domain.User{ID: *userID}
	}

	p.OwnerOrg = o

	return p, nil
}

func (repo *ProjectReposotory) GetByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.ErrForbidden)
	}

	row := repo.QueryRowContext(ctx, qs.ProjectValidDetailByIDQuery, id)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	var (
		p             *domain.Project = &domain.Project{}
		userID, orgID *string
	)
	if err := row.Scan(
		&p.ID, &p.Name, &p.Slug, &p.Description, &p.OwnerType,
		&userID, &orgID,
		&p.IsValid, &p.DeletedAt,
		&p.CreatedAt, &p.UpdatedAt,
	); err != nil {
		return nil, perr.Wrap(err, perr.ErrNotFound)
	}

	// set user
	if userID != nil {
		p.OwnerUser = user
	}

	// set org
	if orgID != nil {
		// @TODO get org detail
		p.OwnerOrg = &domain.Org{
			ID: domain.OrgID(*orgID),
		}
	}

	return p, nil
}

func (repo *ProjectReposotory) Create(ctx context.Context, input domain.ProjectInput) (*string, error) {
	if err := input.Validate(); err != nil {
		return nil, perr.Wrap(err, perr.ErrBadRequest)
	}

	var inputU, slug *string
	ownerType := domain.OwnerTypeUser
	if input.OrgID != nil {
		ownerType = domain.OwnerTypeOrg
		// validate by project capacity
		count, sub, err := repo.CountValidByOrgID(ctx, *input.OrgID)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrBadRequest)
		}
		if err := domain.CanCreateProject(sub, *count, ownerType); err != nil {
			return nil, perr.Wrap(err, perr.ErrBadRequest, err.Error())
		}
	} else {
		cu, err := domain.UserFromCtx(ctx)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrForbidden)
		}
		// validate by project capacity
		count, sub, err := repo.CountValidByUser(ctx, cu.ID)
		if err != nil {
			return nil, perr.Wrap(err, perr.ErrBadRequest)
		}
		if err := domain.CanCreateProject(sub, *count, ownerType); err != nil {
			return nil, perr.Wrap(err, perr.ErrBadRequest, err.Error())
		}

		inputU = tools.StringPtr(cu.ID.String())
	}

	if err := repo.QueryRowContext(
		ctx,
		qs.ProjectCreateQuery,
		input.Name, input.Slug, input.Description, ownerType, inputU, input.OrgID,
	).Scan(&slug); err != nil {
		return nil, perr.Wrap(err, perr.ErrBadRequest)
	}

	return slug, nil
}

func (repo *ProjectReposotory) Delete(ctx context.Context, projectID domain.ProjectID) (int, error) {
	exe, err := repo.ExecContext(ctx, qs.ProjectDeleteQuery, projectID)
	if err != nil {
		return 0, perr.Wrap(err, perr.ErrBadRequest)
	}

	affected, err := exe.RowsAffected()
	if err != nil {
		return 0, perr.Wrap(err, perr.ErrBadRequest)
	}

	return affected, nil
}

func (repo *ProjectReposotory) CountValidByOrgID(ctx context.Context, orgID domain.OrgID) (*int, *domain.Subscription, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, nil, perr.Wrap(err, perr.ErrForbidden)
	}

	row := repo.QueryRowContext(ctx,
		qs.ProjectValidCountByOrgID,
		orgID, cu.ID,
	)
	return countValidByRow(row)
}

func (repo *ProjectReposotory) CountValidByOrgSlug(ctx context.Context, orgSlug string) (*int, *domain.Subscription, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, nil, perr.Wrap(err, perr.ErrForbidden)
	}
	row := repo.QueryRowContext(ctx,
		qs.ProjectValidCountByOrgSlug,
		orgSlug, cu.ID,
	)
	return countValidByRow(row)
}

func (repo *ProjectReposotory) CountValidByUser(ctx context.Context, userID domain.UserID) (*int, *domain.Subscription, error) {
	row := repo.QueryRowContext(ctx,
		qs.ProjectValidCountByUser,
		userID,
	)
	return countValidByRow(row)
}

func countValidByRow(row IRow) (*int, *domain.Subscription, error) {
	if err := row.Err(); err != nil {
		return nil, nil, perr.Wrap(err, perr.ErrNotFound)
	}

	var (
		count                                                              *int
		sID, sSubscriptionID, sCustomerID, sProductID, sSubscriptionStatus *string
		sUserID                                                            *domain.UserID
		sOrgID                                                             *domain.OrgID
		sCreatedAt, sUpdatedAt                                             *time.Time
		s                                                                  *domain.Subscription
	)
	if err := row.Scan(
		&count,
		&sID, &sSubscriptionID, &sCustomerID,
		&sProductID, &sSubscriptionStatus,
		&sUserID, &sOrgID,
		&sCreatedAt, &sUpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return tools.IntPtrAbleZero(0), nil, nil
		}
		return nil, nil, perr.Wrap(err, perr.ErrBadRequest)
	}

	if sID != nil {
		s = &domain.Subscription{
			ID:                       *sID,
			StripeSubscriptionID:     *sSubscriptionID,
			StripeCustomerID:         *sCustomerID,
			StripeProductID:          *sProductID,
			StripeSubscriptionStatus: *sSubscriptionStatus,
			UserID:                   sUserID,
			OrgID:                    sOrgID,
			CreatedAt:                *sCreatedAt,
			UpdatedAt:                *sUpdatedAt,
			IsValid:                  true,
			DeletedAt:                nil,
		}
	}

	return count, s, nil
}
