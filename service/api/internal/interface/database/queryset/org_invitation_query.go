package queryset

const (
	/************************
		List
	*************************/

	// invitation list by org
	// flitered
	OrgInvitationListFromOrgQuery = "SELECT " +
		"r.id, r.status, r.user_type, r.created_at, r.updated_at, " +
		"u.id, u.username, r.email, u.image_url, " +
		"inv.id, inv.username, inv.email, inv.image_url " +
		"FROM org_invitations AS r " +
		"JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true AND o.deleted_at IS NULL " +
		"LEFT JOIN users AS u ON u.id = r.user_id AND u.is_valid = true " + // invited
		"JOIN users AS inv ON inv.id = r.invited_by AND inv.is_valid = true " + // invitor
		// eliminate if current user not belong org
		"JOIN rel_org_members AS rr ON rr.org_id = r.org_id AND rr.user_id = $2 AND rr.is_valid = true " +
		"WHERE r.org_id = $1"

	// invitation list by user
	// filtered by user
	OrgInvitationListQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, " +
		"r.id, r.user_type, r.created_at, " +
		"u.id, u.username, u.email, u.image_url, " + // invitor's information
		"FROM org_invitations AS r " +
		"JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"JOIN users AS u ON u.id = r.invited_by AND u.is_valid = true " +
		"WHERE r.user_id = $1 AND r.status = 'new'"

	// only new
	NewOrgInvitationListQuery = "SELECT " +
		"r.id " +
		"FROM org_invitations AS r " +
		"WHERE r.org_id = $1 AND r.email = $2 AND status = 'new'"

	/************************
		Detail
	*************************/

	// invitation detail
	// filtered by user
	// must filter status = 'new' @controller
	OrgInvitationDetailQuery = "SELECT " +
		"r.id, r.user_type, r.status, r.created_at, " +
		"o.id, o.slug, o.name, o.description, " +
		"u.id, u.username, u.email, u.image_url " + // invitor's information
		"FROM org_invitations AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true AND o.deleted_at IS NULL " +
		"LEFT JOIN users AS u ON u.id = r.invited_by AND u.is_valid = true " +
		"WHERE r.id = $1 AND r.email = $2"

	/************************
		Create
	*************************/

	// invitation insert
	// need filter in repo or con (only ownerType user can invite)
	OrgInvitationCraeteQuery = "INSERT INTO " +
		"org_invitations(org_id, user_id, email, user_type, invited_by) " +
		"VALUES ($1, $2, $3, $4, $5) " +
		"RETURNING id"

	/************************
		Update
	*************************/

	OrgInvitationUpdateStatusQuery = "UPDATE org_invitations " +
		"SET status = $1, updated_at = now() " +
		"WHERE id = $2 AND email = $3 AND status = 'new'"

	/************************
		Delete
	*************************/

	// @TODO delete invitation
)
