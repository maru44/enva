package queryset

const (
	/************************

	invitation

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

	// invitation insert
	// need filter in repo or con (only ownerType user can invite)
	OrgInvitationCraeteQuery = "INSERT INTO " +
		"org_invitations(org_id, user_id, email, user_type, invited_by) " +
		"VALUES ($1, $2, $3, $4, $5) " +
		"RETURNING id"

	OrgInvitationUpdateStatusQuery = "UPDATE org_invitations " +
		"SET status = $1, updated_at = now() " +
		"WHERE id = $2 AND email = $3 AND status = 'new'"

	// @TODO delete invitation

	/************************

	org

	*************************/

	// filtered by user and filtered by valid
	OrgValidListQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.created_at, o.updated_at, " +
		"r.user_type " +
		"FROM rel_org_members AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true AND o.deleted_at IS NULL " +
		"WHERE r.user_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	// filter orgs is_valid on repo or con
	OrgValidDetailBySlugQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.is_valid, o.created_by, o.created_at, o.updated_at, o.deleted_at, " +
		"COUNT(DISTINCT rs.id), r.user_type, " +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status " +
		"FROM orgs AS o " +
		"LEFT JOIN rel_org_members AS rs ON o.id = rs.org_id AND rs.is_valid = true " +
		// eliminate if relation does not exists
		"JOIN rel_org_members AS r ON r.user_id = $1 AND o.id = r.org_id AND r.is_valid = true " +
		"LEFT JOIN subscriptions AS s ON s.org_id = o.id AND s.is_valid = true AND s.deleted_at IS NULL " +
		"WHERE o.slug = $2 AND o.is_valid = true AND o.deleted_at IS NULL " +
		"GROUP BY o.id, r.id, s.id"

	OrgCreateQuery = "INSERT INTO orgs " +
		"(slug, name, description, created_by) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id, slug"

	OrgValidCountByOwner = "SELECT " +
		"COUNT(DISTINCT o.id), " +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status, " +
		"s.user_id, s.org_id, " +
		"s.created_at, s.updated_at " +
		"FROM orgs AS o " +
		"LEFT JOIN subscriptions AS s ON s.user_id = $1 AND s.is_valid = true AND s.deleted_at IS NULL " +
		"WHERE o.created_by = $1 AND o.is_valid = true AND o.deleted_at IS NULL " +
		"GROUP BY s.id"

	/************************

	relation org member

	*************************/

	// also exec after org create
	// must be restricted in repo or con
	// currentUser must be admin or owner
	RelOrgMembersInsertQuery = "INSERT INTO rel_org_members " +
		"(org_id, user_id, user_type, org_invitation_id) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	// member list
	// should be restrict access
	// if this rows do not include current user
	// must throw forbidden error
	OrgUsersQuery = "SELECT " +
		"u.id, u.username, u.email, u.image_url, r.user_type " +
		"FROM rel_org_members AS r " +
		"LEFT JOIN users AS u ON u.id = r.user_id " +
		"WHERE r.org_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	// userType in orgs
	OrgUserTypeQuery = "SELECT " +
		"r.user_type " +
		"FROM rel_org_members AS r " +
		"WHERE r.org_id = $1 AND r.user_id = $2 AND r.is_valid = true AND r.deleted_at IS NULL"

	OrgUserTypeByEmailQuery = "SELECT " +
		"r.user_type " +
		"FROM users AS u " +
		"JOIN rel_org_members AS r ON r.org_id = $1 AND r.user_id = u.id AND r.is_valid = true AND r.deleted_at IS NULL " +
		"JOIN orgs AS o ON o.id = $1 AND o.is_valid = true AND o.deleted_at IS NULL " +
		"WHERE u.email = $2"

	// if updated user's origin user type is 'owner'
	// validate is current user 'owner'
	OrgEliminateMemberQuery = "UPDATE rel_org_members " +
		"SET is_valid = false, updated_at = now(), deleted_at = now() " +
		"WHERE org_id = $1 AND user_id = $2"

	OrgMemberUserTypeUpdateQuery = "UPDATE rel_org_members " +
		"SET user_type = $1, updated_at = now() " +
		"WHERE org_id = $2 AND user_id = $3"

	OrgReAddMemberQuery = ""

	// filter by access user
	// is it needed?
	// OrgMemberCountByOrgID = "SELECT " +
	// 	"COUNT(DISTINCT rs.id), s.stripe_product_id " +
	// 	"FROM rel_org_members AS rs " +
	// 	"LEFT JOIN subscriptions AS s ON s.org_id = $1 AND s.is_valid = true AND s.deleted_at IS NULL " +
	// 	"JOIN rel_org_members AS r ON r.user_id = $2 AND r.org_id = $1 AND r.is_valid = true " +
	// 	"WHERE rs.org_id = $1 AND rs.is_valid = true AND rs.deleted_at IS NULL"
)
