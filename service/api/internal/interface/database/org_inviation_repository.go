package database

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/internal/interface/mysmtp"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgInvitationRepository struct {
	ISqlHandler
	mysmtp.ISmtpHandler
}

func (repo *OrgInvitationRepository) ListFromOrg(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitation, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	rows, err := repo.QueryContext(ctx,
		queryset.OrgInvitationListFromOrgQuery,
		orgID, cu.ID,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var orgs []domain.OrgInvitation
	for rows.Next() {
		var (
			o   domain.OrgInvitation
			u   domain.User
			inv domain.User
		)
		if err := rows.Scan(
			&o.ID, &o.Status, &o.UserType, &o.CreatedAt, &o.UpdatedAt,
			&u.ID, &u.Username, &u.Email, &u.ImageURL,
			&inv.ID, &inv.Username, &inv.Email, &inv.ImageURL,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}
		o.Org = domain.Org{ID: orgID}
		o.User = u
		o.Invitor = inv
		orgs = append(orgs, o)
	}

	return orgs, nil
}

func (repo *OrgInvitationRepository) List(ctx context.Context) ([]domain.OrgInvitation, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.OrgInvitationListQuery,
		cu.ID,
	)
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var orgs []domain.OrgInvitation
	for rows.Next() {
		var (
			o   domain.OrgInvitation
			inv domain.User
			org domain.Org
		)
		if err := rows.Scan(
			&org.ID, &org.Slug, &org.Name, &org.Description,
			&o.ID, &o.UserType, &o.CreatedAt,
			&inv.ID, &inv.Username, &inv.Email, &inv.ImageURL,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}
		o.Org = org
		o.User = *cu
		o.Invitor = inv
		orgs = append(orgs, o)
	}

	return orgs, nil
}

func (repo *OrgInvitationRepository) Detail(ctx context.Context, invID domain.OrgInvitationID) (*domain.OrgInvitation, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	row := repo.QueryRowContext(ctx,
		queryset.OrgInvitationDetailQuery,
		invID, cu.ID,
	)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var (
		o   *domain.OrgInvitation
		org domain.Org
		inv domain.User
	)
	if err := row.Scan(
		&o.ID, &o.UserType, &o.Status, &o.CreatedAt,
		&org.ID, &org.Slug, &org.Name, &org.Description,
		&inv.ID, &inv.Username, &inv.Email, &inv.ImageURL,
	); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	o.Org = org
	o.Invitor = inv
	o.User = *cu

	return o, nil
}

func (repo *OrgInvitationRepository) Create(ctx context.Context, input domain.OrgInvitationInput, targetID domain.UserID) error {
	if err := input.Validate(); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	var id *string
	if err := repo.QueryRowContext(ctx,
		queryset.OrgInvitationCraeteQuery,
		input.OrgID, input.UserID, input.Eamil, input.UserType, cu.ID,
	).Scan(&id); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	mailInput := domain.SmtpInput{
		Subject: "",
		Message: "",
		To:      input.Eamil,
	}
	if err := repo.Send(mailInput); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}

func (repo *OrgInvitationRepository) ListPastInvitations(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitationID, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.PastOrgInvitationListQuery,
		orgID, cu.Email,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var ids []domain.OrgInvitationID
	for rows.Next() {
		var id domain.OrgInvitationID
		if err := rows.Scan(
			&id,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (repo *OrgInvitationRepository) UpdateStatus(ctx context.Context, invID domain.OrgInvitationID, status domain.OrgInvitationStatus) error {
	res, err := repo.ExecContext(ctx,
		queryset.OrgInvitationUpdateStatusQuery,
		status, invID,
	)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	if affected == 0 {
		return perr.New("no rows affected", perr.BadRequest)
	}

	return nil
}
