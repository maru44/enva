package queryset

const (
	OrgListQuery = "SELECT " +
		"id, slug, name, description, created_by, created_at, updated_at " +
		"LEFT JOIN rel_org_members AS users ON orgs.id = users.org_id " +
		"FROM orgs " +
		"WHERE user_id = ?"

	OrgsQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.created_by, o.created_at, o.updated_at " +
		"FROM rel_org_members AS r " +
		"LEFT JOIN orgs AS o ON o.id = r.org_id AND o.is_valid = true " +
		"WHERE is_valid = true AND deleted_at IS NULL AND user_id = ?"

		// @TODO filter is_valid on repo
	OrgDetailQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.is_valid, o.created_by, o.created_at, o.updated_at, " +
		"u.id, u.email, u.username, u.image_url " +
		"JOIN rel_org_members AS r ON r.org_id = AND r.is_valid = true " +
		"JOIN users AS u ON u.id = r.user_id " +
		"FROM orgs AS o " +
		"WHERE o.id = ?"

	OrgDetailBySlugQuery = "SELECT " +
		"o.id, o.slug, o.name, o.description, o.is_valid, o.created_by, o.created_at, o.updated_at, " +
		"u.id, u.email, u.username, u.image_url " +
		"JOIN rel_org_members AS r ON r.org_id = AND r.is_valid = true " +
		"JOIN users AS u ON u.id = r.user_id " +
		"FROM orgs AS o " +
		"WHERE o.slug = ?"

	OrgCreateQuery = ""
)
