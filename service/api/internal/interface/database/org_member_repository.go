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

	// current user
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.NotFound)
	}
	if cu.ID != input.UserID {
		return perr.New("current user is not invited user", perr.Forbidden)
	}

	tx, err := repo.BeginTx(ctx, nil)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	// create member
	var id *string
	if err := tx.QueryRowContext(ctx,
		queryset.RelOrgMembersInsertQuery,
		input.OrgID, input.UserID, input.UserType, input.OrgInvitationID,
	).Scan(&id); err != nil {
		tx.Rollback()
		return perr.Wrap(err, perr.BadRequest)
	}

	// update invitation status
	res, err := tx.ExecContext(ctx,
		queryset.OrgInvitationUpdateStatusQuery,
		domain.OrgInvitationStatusAccepted,
		input.OrgInvitationID,
	)
	if err != nil {
		tx.Rollback()
		return perr.Wrap(err, perr.BadRequest)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return perr.Wrap(err, perr.BadRequest)
	}
	if affected == 0 {
		tx.Rollback()
		return perr.New("no rows affected", perr.BadRequest)
	}

	// past invitation ids
	rows, err := repo.QueryContext(ctx,
		queryset.PastOrgInvitationListQuery,
		input.OrgID, cu.Email,
	)
	if err != nil {
		tx.Rollback()
		return perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		tx.Rollback()
		return perr.Wrap(err, perr.NotFound)
	}

	var pastIDs []domain.OrgInvitationID
	for rows.Next() {
		var id domain.OrgInvitationID
		if err := rows.Scan(
			&id,
		); err != nil {
			tx.Rollback()
			return perr.Wrap(err, perr.NotFound)
		}
		pastIDs = append(pastIDs, id)
	}

	// update past invitations' status >> closed
	for _, invID := range pastIDs {
		res, err := repo.ExecContext(ctx,
			queryset.OrgInvitationUpdateStatusQuery,
			domain.OrgInvitationStatusClosed, invID,
		)
		if err != nil {
			tx.Rollback()
			return perr.Wrap(err, perr.BadRequest)
		}

		affected, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			return perr.Wrap(err, perr.BadRequest)
		}
		if affected == 0 {
			tx.Rollback()
			return perr.New("no rows affected", perr.BadRequest)
		}
	}

	tx.Commit()

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

func (repo *OrgMemberRepository) GetCurrentUserType(ctx context.Context, orgID domain.OrgID) (*domain.UserType, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	row := repo.QueryRowContext(ctx, queryset.OrgUserTypeQuery, orgID, user.ID)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}
	var ut *domain.UserType
	if err := row.Scan(&ut); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return ut, nil
}
