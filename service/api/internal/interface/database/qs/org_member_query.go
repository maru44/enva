package qs

const (
	/************************
		list
	*************************/
	// member list
	// should be restrict access
	// if this rows do not include current user
	// must throw forbidden error
	OrgUsersQuery = "SELECT " +
		"u.id, u.username, u.email, u.image_url, r.user_type " +
		"FROM rel_org_members AS r " +
		"LEFT JOIN users AS u ON u.id = r.user_id " +
		"WHERE r.org_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	/************************
		Delete
	*************************/

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

	/************************
		Create
	*************************/

	// also exec after org create
	// must be restricted in repo or con
	// currentUser must be admin or owner
	RelOrgMembersInsertQuery = "INSERT INTO rel_org_members " +
		"(org_id, user_id, user_type, org_invitation_id) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	/************************
		Delete
	*************************/

	// if updated user's origin user type is 'owner'
	// validate is current user 'owner'
	OrgEliminateMemberQuery = "UPDATE rel_org_members " +
		"SET is_valid = false, updated_at = now(), deleted_at = now() " +
		"WHERE org_id = $1 AND user_id = $2"

	/************************
		Update
	*************************/

	OrgMemberUserTypeUpdateQuery = "UPDATE rel_org_members " +
		"SET user_type = $1, updated_at = now() " +
		"WHERE org_id = $2 AND user_id = $3"

	/************************
		Count
	*************************/

	// filter by access user
	OrgMemberCountByOrgID = "SELECT " +
		"COUNT(DISTINCT rs.id), " +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status, " +
		"s.user_id, s.org_id, " +
		"s.created_at, s.updated_at " +
		"FROM rel_org_members AS rs " +
		"LEFT JOIN subscriptions AS s ON s.org_id = $1 AND s.is_valid = true AND s.deleted_at IS NULL " +
		"JOIN rel_org_members AS r ON r.user_id = $2 AND r.org_id = $1 AND r.is_valid = true " +
		"WHERE rs.org_id = $1 AND rs.is_valid = true AND rs.deleted_at IS NULL " +
		"GROUP BY s.id"
)
