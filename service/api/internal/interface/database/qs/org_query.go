package qs

const (
	/************************
		List
	*************************/

	// filtered by user and filtered by valid
	OrgValidListQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.created_at, o.updated_at, " +
		"r.user_type " +
		"FROM rel_org_members AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true AND o.deleted_at IS NULL " +
		"WHERE r.user_id = $1 AND r.is_valid = true AND r.deleted_at IS NULL"

	/************************
		Detail
	*************************/

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

	/************************
		Create
	*************************/

	OrgCreateQuery = "INSERT INTO orgs " +
		"(slug, name, description, created_by) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id, slug"

	/************************
		Count
	*************************/

	OrgValidCountByOwner = "SELECT " +
		"COUNT(DISTINCT o.id), " +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status, " +
		"s.user_id, s.org_id, " +
		"s.created_at, s.updated_at " +
		"FROM orgs AS o " +
		"LEFT JOIN subscriptions AS s ON s.user_id = $1 AND s.is_valid = true AND s.deleted_at IS NULL " +
		"WHERE o.created_by = $1 AND o.is_valid = true AND o.deleted_at IS NULL " +
		"GROUP BY s.id " +
		"FOR UPDATE"
)
