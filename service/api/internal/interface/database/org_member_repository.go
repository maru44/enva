package database

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgMemberRepository struct {
	ISqlHandler
}

func (repo *OrgMemberRepository) Create(ctx context.Context, input domain.OrgMemberInput) error {
	if err := input.Validate(ctx); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	var id *string
	if err := repo.QueryRowContext(ctx,
		queryset.RelOrgMembersInsertQuery,
		input.OrgID, input.UserID, input.UserType, input.OrgInvitationID,
	).Scan(&id); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}

func (repo *OrgMemberRepository) List(ctx context.Context, orgID domain.OrgID) (map[domain.UserType][]domain.User, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	// confirm access
	var userType domain.UserType
	row := repo.QueryRowContext(ctx,
		queryset.OrgUserTypeQuery,
		orgID, cu.ID,
	)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}
	if err := row.Scan(&userType); err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	} else if userType == domain.UserType("") {
		return nil, perr.New("Not belong to this org", perr.Forbidden)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.OrgUsersQuery,
		orgID,
	)
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	var users map[domain.UserType][]domain.User
	for rows.Next() {
		var (
			u  domain.User
			ut domain.UserType
		)
		if err := rows.Scan(
			&u.ID, &u.Username, &u.Email, &u.ImageURL, &ut,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}

		users[ut] = append(users[ut], u)
	}

	return users, nil
}
