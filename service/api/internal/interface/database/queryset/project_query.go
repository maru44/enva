package queryset

const (
	/*********************************
		List
	*********************************/

	ProjectValidListByUserQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at IS NULL AND owner_user_id = $1"

	ProjectValidSlugListByUserQuery = "SELECT " +
		"slug " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at IS NULL AND owner_user_id = $1"

	// list by organization
	ProjectListValidByOrgQuery = "SELECT " +
		"p.id, p.name, p.slug, p.description, p.owner_type, p.owner_user_id, p.owner_org_id, " +
		"p.created_at, p.updated_at " +
		"FROM projects AS p " +
		// filter current user is member of org
		"JOIN rel_org_members AS r ON r.org_id = $1 AND r.user_id = $2 AND r.is_valid = true " +
		"WHERE p.is_valid = true AND p.deleted_at IS NULL AND p.owner_org_id = $1"

	/*********************************
		Detail
	*********************************/

	// only by user
	ProjectValidDetailBySlugQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, deleted_at, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE slug = $1 AND owner_user_id = $2 AND is_valid = true AND deleted_at IS NULL"

	ProjectValidDetailBySlugAndOrgIdQuery = "SELECT " +
		"p.id, p.name, p.slug, p.description, p.owner_type, p.owner_user_id, " +
		"p.is_valid, p.deleted_at, " +
		"p.created_at, p.updated_at, " +
		"o.slug, o.name " +
		"FROM projects AS p " +
		"JOIN rel_org_members AS r ON r.user_id = $1 AND r.org_id = $2 AND r.is_valid = true AND r.deleted_at IS NULL " +
		"JOIN orgs AS o ON o.id = $2 AND o.is_valid = true " +
		"WHERE p.slug = $3 AND p.owner_org_id = $2 AND is_valid = true AND p.deleted_at IS NULL"

	ProjectValidDetailBySlugAndOrgSlugQuery = "SELECT " +
		"p.id, p.name, p.slug, p.description, p.owner_type, p.owner_user_id, " +
		"p.is_valid, p.deleted_at, " +
		"p.created_at, p.updated_at, " +
		"o.id, o.name " +
		"FROM orgs AS o " +
		"JOIN projects AS p ON p.slug = $3 AND p.owner_org_id = o.id AND p.is_valid = true AND p.deleted_at IS NULL " +
		"JOIN rel_org_members AS r ON r.user_id = $1 AND r.org_id = o.id AND r.is_valid = true AND r.deleted_at IS NULL " +
		"WHERE o.is_valid = true AND o.slug= $2"

	ProjectValidDetailByIDQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, deleted_at, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE id = $1 AND is_valid = true AND deleted_at IS NULL"

	/*********************************
		Create
	*********************************/

	ProjectCreateQuery = "INSERT INTO " +
		"projects(name, slug, description, owner_type, owner_user_id, owner_org_id) " +
		"VALUES($1, $2, $3, $4, $5, $6) " +
		"RETURNING slug"

	/*********************************
		Update
	*********************************/

	// @TODO update

	/*********************************
		Delete
	*********************************/

	ProjectDeleteQuery = "UPDATE projects " +
		"SET deleted_at = now(), is_valid = false, updated_at = now() " +
		"WHERE id = $1"

	/*********************************
		Count
	*********************************/

	ProjectValidCountByOrgID = "SELECT " +
		"COUNT(DISTINCT p.id)," +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status, " +
		"s.user_id, s.org_id, " +
		"s.created_at, s.updated_at " +
		"FROM orgs AS o " +
		"LEFT JOIN projects AS p ON p.owner_org_id = o.id AND p.owner_type = 'org' AND p.is_valid = true AND p.deleted_at IS NULL " +
		"LEFT JOIN subscriptions AS s ON s.org_id = o.id AND s.is_valid = true AND s.deleted_at IS NULL " +
		// validate current user access
		"JOIN rel_org_members AS r ON r.org_id = $1 AND r.user_id = $2 AND r.is_valid = true AND r.deleted_at IS NULL " +
		"WHERE o.id = $1 AND o.is_valid = true AND o.deleted_at IS NULL " +
		"GROUP BY s.id"

	ProjectValidCountByOrgSlug = "SELECT " +
		"COUNT(DISTINCT p.id)," +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status, " +
		"s.user_id, s.org_id, " +
		"s.created_at, s.updated_at " +
		"FROM orgs AS o " +
		"LEFT JOIN projects AS p ON p.owner_org_id = o.id AND p.owner_type = 'org' AND p.is_valid = true AND p.deleted_at IS NULL " +
		"LEFT JOIN subscriptions AS s ON s.org_id = o.id AND s.is_valid = true AND s.deleted_at IS NULL " +
		// validate current user access
		"JOIN rel_org_members AS r ON r.org_id = o.id AND r.user_id = $2 AND r.is_valid = true AND r.deleted_at IS NULL " +
		"WHERE o.slug = $1 AND o.is_valid = true AND o.deleted_at IS NULL " +
		"GROUP BY o.id, s.id"

	ProjectValidCountByUser = "SELECT " +
		"COUNT(DISTINCT p.id), " +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status, " +
		"s.user_id, s.org_id, " +
		"s.created_at, s.updated_at " +
		"FROM projects AS p " +
		"LEFT JOIN subscriptions AS s ON s.user_id = $1 AND s.is_valid = true AND s.deleted_at IS NULL " +
		"WHERE p.owner_user_id = $1 AND p.owner_type = 'user' AND p.is_valid = true AND p.deleted_at IS NULL " +
		"GROUP BY s.id"
)
