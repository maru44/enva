package queryset

const (
	/************************

	invitation

	*************************/

	// invitation list by user
	OrgInvitationListQuery = "SELECT " +
		"o.id, o.slug, o.name, r.created_at, u.username " +
		"FROM org_invitations AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"LEFT JOIN users AS u ON u.id = r.invited_by AND u.is_valid = true " +
		"WHERE r.user_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	// invitation detail
	OrgInvitationDetailQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, " +
		"r.created_at, u.username " +
		"FROM org_invitations AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"LEFT JOIN users AS u ON u.id = r.invited_by AND u.is_valid = true " +
		"WHERE r.id = $1"

	// invitation insert
	OrgInvitationCraeteQuery = "INSERT org_invitations INTO " +
		"(org_id, user_id, user_type, invited_by) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	/************************

	org

	*************************/

	OrgListQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.created_by, o.created_at, o.updated_at, " +
		"r.user_type " +
		"FROM rel_org_members AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"WHERE r.user_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	// @TODO filter orgs is_valid on repo or con
	OrgDetailQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.is_valid, o.created_by, o.created_at, o.updated_at, " +
		"r.user_type " +
		"u.id, u.email, u.username, u.image_url " +
		"JOIN rel_org_members AS r ON r.org_id = AND r.is_valid = true AND r.deleted_at IS NULL " +
		"JOIN users AS u ON u.id = r.user_id AND u.is_valid = true AND u.is_email_verified = true " +
		"FROM orgs AS o " +
		"WHERE o.id = ?"

	// @TODO filter orgs is_valid on repo or con
	OrgDetailBySlugQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.is_valid, o.created_by, o.created_at, o.updated_at, " +
		"r.user_type " +
		"u.id, u.email, u.username, u.image_url " +
		"JOIN rel_org_members AS r ON r.org_id = AND r.is_valid = true AND r.deleted_at IS NULL " +
		"JOIN users AS u ON u.id = r.user_id AND u.is_valid = true AND u.is_email_verified = true " +
		"FROM orgs AS o " +
		"WHERE o.slug = ?"

	OrgCreateQuery = "INSERT INTO orgs " +
		"(slug, name, description, created_by) " +
		"VALUES ($1, $2, $3, $4)"

	/************************

	relation org member

	*************************/

	// also exec after org create
	RelOrgMembersInsertQuery = "INSERT INTO rel_org_members " +
		"(org_id, user_id, user_type, org_invitation_id) " +
		"VALUES ($1, $2, $3, $4)"

	OrgEliminateMemberQuery = "UPDATE rel_org_members " +
		"SET is_valid = true AND deleted_at = now()" +
		"WHERE org_id = $1 AND user_id = $2"

	OrgReAddMemberQuery = ""
)
