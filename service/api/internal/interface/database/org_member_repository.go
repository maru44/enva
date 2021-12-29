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
