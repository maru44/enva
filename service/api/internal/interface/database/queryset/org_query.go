package queryset

const (
	/************************

	invitation

	*************************/

	// invitation list by org
	OrgInvitationListFromOrgQuery = "SELECT " +
		"r.id, r.is_valid, r.user_type, r.created_at, r.updated_at, r.deleted_at " +
		"u.id, u.username, u.email, u.image_url, " +
		"inv.id, inv.username, inv.email, inv.image_url, " +
		"FROM org_invitations AS r " +
		"LEFT JOIN orgs AS o ON o.id = $1 AND o.is_valid = true " +
		"JOIN users AS u ON u.id = r.user_id AND u.is_valid = true " + // invited
		"JOIN users AS inv ON inv.id = r.invited_by AND inv.is_valid = true " + // invitor
		// eliminate if current user not belong org
		"JOIN rel_org_members AS rr ON rr.org_id = $1 AND rr.user_id = $2 AND rr.is_valid = true " +
		"WHERE r.org_id = $1"

	// invitation list by user
	OrgInvitationListQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, " +
		"r.id, r.user_type, r.created_at, " +
		"u.id, u.username, u.email, u.image_url, " + // invitor's information
		"FROM org_invitations AS r " +
		"JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"JOIN users AS u ON u.id = r.invited_by AND u.is_valid = true " +
		"WHERE r.user_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	// invitation detail
	OrgInvitationDetailQuery = "SELECT " +
		"r.id, r.user_type, r.is_valid, r.created_at, " +
		"o.id, o.slug, o.name, o.description, " +
		"u.id, u.username, u.email, u.image_url " + // invitor's information
		"FROM org_invitations AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"LEFT JOIN users AS u ON u.id = r.invited_by AND u.is_valid = true " +
		"WHERE r.id = $1 AND r.user_id = $2"

	// invitation insert
	OrgInvitationCraeteQuery = "INSERT org_invitations INTO " +
		"(org_id, user_id, user_type, invited_by) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	/************************

	org

	*************************/

	OrgListQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.created_at, o.updated_at, " +
		"r.user_type " +
		"FROM rel_org_members AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"WHERE r.user_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	// @TODO filter orgs is_valid on repo or con
	OrgDetailQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.is_valid, o.created_by, o.created_at, o.updated_at, " +
		"COUNT(DISTINCT rs.id) " +
		"LEFT JOIN rel_org_members AS rs ON o.id = rs.org_id AND rs.is_valid = true " +
		// eliminate if relation does not exists
		"JOIN rel_org_members AS r ON o.user_id = $1 AND o.id = r.org_id AND r.is_valid = true " +
		"FROM orgs AS o " +
		"WHERE o.id = $2"

	// @TODO filter orgs is_valid on repo or con
	OrgDetailBySlugQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.is_valid, o.created_by, o.created_at, o.updated_at, " +
		"COUNT(DISTINCT r.id) " +
		"LEFT JOIN rel_org_members AS rs ON o.id = rs.org_id AND rs.is_valid = true " +
		// eliminate if relation does not exists
		"JOIN rel_org_members AS r ON o.user_id = $1 AND o.id = r.org_id AND r.is_valid = true " +
		"FROM orgs AS o " +
		"WHERE o.slug = $1"

	OrgUsersQuery = "SELECT " +
		"u.id, u.username, u.email, u.image_url, r.user_type " +
		"LEFT JOIN users AS u ON u.id = r.user_id AND " +
		"FROM rel_org_members AS r " +
		"WHERE r.org_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	OrgCreateQuery = "INSERT INTO orgs " +
		"(slug, name, description, created_by) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id, slug"

	/************************

	relation org member

	*************************/

	// also exec after org create
	RelOrgMembersInsertQuery = "INSERT INTO rel_org_members " +
		"(org_id, user_id, user_type, org_invitation_id) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	OrgEliminateMemberQuery = "UPDATE rel_org_members " +
		"SET is_valid = true AND deleted_at = now()" +
		"WHERE org_id = $1 AND user_id = $2"

	OrgReAddMemberQuery = ""
)
