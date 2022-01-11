package database

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database/queryset"
	"github.com/maru44/enva/service/api/internal/interface/mysmtp"
	"github.com/maru44/enva/service/api/pkg/domain"
	"github.com/maru44/perr"
)

type OrgRepository struct {
	ISqlHandler
	mysmtp.ISmtpHandler
}

func (repo *OrgRepository) List(ctx context.Context) ([]domain.Org, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.OrgValidListQuery, user.ID,
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
		queryset.OrgValidListQuery, user.ID,
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

func (repo *OrgRepository) DetailBySlug(ctx context.Context, slug string) (*domain.Org, *domain.UserType, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}
	row := repo.QueryRowContext(
		ctx,
		queryset.OrgValidDetailBySlugQuery,
		user.ID, slug,
	)
	if err := row.Err(); err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}
	var (
		o                                                                  *domain.Org = &domain.Org{}
		u                                                                  domain.User
		ut                                                                 *domain.UserType
		sID, sSubscriptionID, sCustomerID, sProductID, sSubscriptionStatus *string
	)
	if err := row.Scan(
		&o.ID, &o.Slug, &o.Name, &o.Description, &o.IsValid,
		&u.ID, &o.CreatedAt, &o.UpdatedAt, &o.DeletedAt, &o.UserCount,
		&ut,
		&sID, &sSubscriptionID, &sCustomerID, &sProductID, &sSubscriptionStatus,
	); err != nil {
		return nil, nil, perr.Wrap(err, perr.NotFound)
	}

	if sID != nil {
		o.Subscription = &domain.Subscription{
			ID:                       *sID,
			StripeSubscriptionID:     *sSubscriptionID,
			StripeCustomerID:         *sCustomerID,
			StripeProductID:          *sProductID,
			StripeSubscriptionStatus: *sSubscriptionStatus,
		}
	}

	o.CreatedBy = u
	return o, ut, nil
}

func (repo *OrgRepository) Create(ctx context.Context, input domain.OrgInput) (*string, error) {
	if err := input.Validate(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	// validate count org
	count, sub, err := repo.OrgValidCount(ctx, cu.ID)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}
	if err := domain.CanCreateOrg(sub, *count); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest, err.Error())
	}

	tx, err := repo.BeginTx(ctx, nil)
	if err != nil {
		return nil, perr.Wrap(err, perr.InternalServerError)
	}

	var id, slug *string
	if err := tx.QueryRowContext(ctx,
		queryset.OrgCreateQuery,
		input.Slug, input.Name, input.Description, cu.ID,
	).Scan(&id, &slug); err != nil {
		tx.Rollback()
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	var memberID *string
	if err := tx.QueryRowContext(ctx,
		queryset.RelOrgMembersInsertQuery,
		id, cu.ID, domain.UserTypeOwner, nil,
	).Scan(&memberID); err != nil {
		tx.Rollback()
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	tx.Commit()

	return slug, nil
}

func (repo *OrgRepository) OrgValidCount(ctx context.Context, userID domain.UserID) (*int, *domain.Subscription, error) {
	row := repo.QueryRowContext(ctx,
		queryset.OrgValidCountByOwner, userID,
	)
	return countValidByRow(row)
}

/*************************

invitation

*************************/

func (repo *OrgRepository) InvitationListFromOrg(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitation, error) {
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

	var invs []domain.OrgInvitation
	for rows.Next() {
		var (
			id, username, email, imageUrl *string
			o                             domain.OrgInvitation
			u                             domain.User
			inv                           domain.User
		)
		if err := rows.Scan(
			&o.ID, &o.Status, &o.UserType, &o.CreatedAt, &o.UpdatedAt,
			&id, &username, &email, &imageUrl,
			&inv.ID, &inv.Username, &inv.Email, &inv.ImageURL,
		); err != nil {
			return nil, perr.Wrap(err, perr.NotFound)
		}
		u = domain.User{Email: *email}
		if id != nil {
			u = domain.User{
				ID:       domain.UserID(*id),
				Username: *username,
				Email:    *email,
				ImageURL: imageUrl,
			}
		}
		o.Org = domain.Org{ID: orgID}
		o.User = u
		o.Invitor = inv
		invs = append(invs, o)
	}

	return invs, nil
}

func (repo *OrgRepository) InvitationList(ctx context.Context) ([]domain.OrgInvitation, error) {
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

func (repo *OrgRepository) InvitationDetail(ctx context.Context, invID domain.OrgInvitationID) (*domain.OrgInvitation, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	row := repo.QueryRowContext(ctx,
		queryset.OrgInvitationDetailQuery,
		invID, cu.Email,
	)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	var (
		o   *domain.OrgInvitation = &domain.OrgInvitation{}
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

func (repo *OrgRepository) Invite(ctx context.Context, input domain.OrgInvitationInput) error {
	if err := input.Validate(); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	// if user have input's email exists already
	if _, err := repo.memberGetUserTypeByEmail(ctx, input.OrgID, input.Eamil); err == nil {
		errStr := "user already belongs to " + input.OrgName
		return perr.New(errStr, perr.BadRequest, errStr)
	}

	count, sub, err := repo.MemberValidCount(ctx, input.OrgID)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	if err := domain.CanCreateOrgMember(sub, *count); err != nil {
		return perr.Wrap(err, perr.BadRequest, err.Error())
	}

	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	var uId *domain.UserID
	if input.User != nil {
		uId = &input.User.ID
	}

	var id *string
	if err := repo.QueryRowContext(ctx,
		queryset.OrgInvitationCraeteQuery,
		input.OrgID, uId, input.Eamil, input.UserType, cu.ID,
	).Scan(&id); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	// send invitation mail
	mailInput := input.CreateMail(domain.OrgInvitationID(*id), *cu)
	if err := repo.Send(mailInput); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	return nil
}

func (repo *OrgRepository) InvitationPastList(ctx context.Context, orgID domain.OrgID) ([]domain.OrgInvitationID, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.NewOrgInvitationListQuery,
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

func (repo *OrgRepository) InvitationUpdateStatus(ctx context.Context, invID domain.OrgInvitationID, status domain.OrgInvitationStatus) error {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}

	res, err := repo.ExecContext(ctx,
		queryset.OrgInvitationUpdateStatusQuery,
		status, invID, cu.Email,
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

func (repo *OrgRepository) InvitationDeny(ctx context.Context, invID domain.OrgInvitationID) error {
	inv, err := repo.InvitationDetail(ctx, invID)
	if err != nil {
		return perr.Wrap(err, perr.NotFound)
	}
	if err := repo.InvitationUpdateStatus(ctx, invID, domain.OrgInvitationStatusDenied); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	ids, err := repo.InvitationPastList(ctx, inv.Org.ID)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	for _, id := range ids {
		if err := repo.InvitationUpdateStatus(ctx, id, domain.OrgInvitationStatusClosed); err != nil {
			return perr.Wrap(err, perr.BadRequest)
		}
	}

	return nil
}

/*************************

member

*************************/

func (repo *OrgRepository) MemberCreate(ctx context.Context, input domain.OrgMemberInput) error {
	if err := input.Validate(ctx); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	count, sub, err := repo.MemberValidCount(ctx, input.OrgID)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	if err := domain.CanCreateOrgMember(sub, *count); err != nil {
		return perr.Wrap(err, perr.BadRequest, err.Error())
	}

	// current user
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.NotFound)
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
		input.OrgInvitationID, cu.Email,
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
	rows, err := tx.QueryContext(ctx,
		queryset.NewOrgInvitationListQuery,
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
		res, err := tx.ExecContext(ctx,
			queryset.OrgInvitationUpdateStatusQuery,
			domain.OrgInvitationStatusClosed, invID, cu.Email,
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

func (repo *OrgRepository) MemberList(ctx context.Context, orgID domain.OrgID) (map[domain.UserType][]domain.User, error) {
	// confirm access
	if _, err := repo.MemberGetCurrentUserType(ctx, orgID); err != nil {
		return nil, perr.Wrap(err, perr.Forbidden)
	}

	rows, err := repo.QueryContext(ctx,
		queryset.OrgUsersQuery,
		orgID,
	)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}
	if err := rows.Err(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	users := map[domain.UserType][]domain.User{}
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

func (repo *OrgRepository) MemberGetCurrentUserType(ctx context.Context, orgID domain.OrgID) (*domain.UserType, error) {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return repo.MemberGetUserType(ctx, user.ID, orgID)
}

func (repo *OrgRepository) MemberGetUserType(ctx context.Context, userID domain.UserID, orgID domain.OrgID) (*domain.UserType, error) {
	row := repo.QueryRowContext(ctx, queryset.OrgUserTypeQuery, orgID, userID)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}
	var ut *domain.UserType
	if err := row.Scan(&ut); err != nil {
		return nil, perr.Wrap(err, perr.BadRequest)
	}

	return ut, nil
}

func (repo *OrgRepository) MemberUpdateUserType(ctx context.Context, input domain.OrgMemberUpdateInput) error {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	currentUt, err := repo.MemberGetCurrentUserType(ctx, input.OrgID)
	if err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}
	if err := currentUt.IsAdmin(); err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}

	if err := input.Validate(); err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	if user.ID == input.UserID {
		return perr.New("cannot change own user type", perr.Forbidden, "cannot change your user type by yourself")
	}
	updatedUserUt, err := repo.MemberGetUserType(ctx, input.UserID, input.OrgID)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}

	// if target or new value is owner, current user must be owner
	if *updatedUserUt == domain.UserTypeOwner || *input.UserType == domain.UserTypeOwner {
		if *currentUt != domain.UserTypeOwner {
			return perr.New("User is not owner", perr.Forbidden, "you are not owner")
		}
	}

	res, err := repo.ExecContext(ctx,
		queryset.OrgMemberUserTypeUpdateQuery,
		input.UserType, input.OrgID, input.UserID,
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

func (repo *OrgRepository) MemberDelete(ctx context.Context, userID domain.UserID, orgID domain.OrgID) error {
	user, err := domain.UserFromCtx(ctx)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	if user.ID == userID {
		return perr.New("cannot delete own", perr.Forbidden, "cannot delete yourself")
	}

	currentUt, err := repo.MemberGetCurrentUserType(ctx, orgID)
	if err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}
	if err := currentUt.IsAdmin(); err != nil {
		return perr.Wrap(err, perr.Forbidden)
	}

	updatedUserUt, err := repo.MemberGetUserType(ctx, userID, orgID)
	if err != nil {
		return perr.Wrap(err, perr.BadRequest)
	}
	// only owner user can delete owner
	if *updatedUserUt == domain.UserTypeOwner && *currentUt != domain.UserTypeOwner {
		return perr.New("User is not owner", perr.Forbidden, "you are not owner")
	}

	res, err := repo.ExecContext(ctx,
		queryset.OrgEliminateMemberQuery,
		orgID, userID,
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

func (repo *OrgRepository) MemberValidCount(ctx context.Context, orgID domain.OrgID) (*int, *domain.Subscription, error) {
	cu, err := domain.UserFromCtx(ctx)
	if err != nil {
		return nil, nil, perr.Wrap(err, perr.BadRequest)
	}
	row := repo.QueryRowContext(ctx,
		queryset.OrgMemberCountByOrgID,
		orgID, cu.ID,
	)
	return countValidByRow(row)
}

/*********************

util

*********************/

func (repo *OrgRepository) memberGetUserTypeByEmail(ctx context.Context, orgID domain.OrgID, email string) (*domain.UserType, error) {
	row := repo.QueryRowContext(ctx,
		queryset.OrgUserTypeByEmailQuery,
		orgID, email,
	)
	if err := row.Err(); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}
	var ut *domain.UserType
	if err := row.Scan(&ut); err != nil {
		return nil, perr.Wrap(err, perr.NotFound)
	}

	return ut, nil
}
