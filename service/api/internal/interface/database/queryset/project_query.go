package queryset

const (
	ProjectListQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at IS NULL"

	ProjectListByUserQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at IS NULL AND owner_user_id = $1"

	ProjectSlugListByUserQuery = "SELECT " +
		"slug " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at IS NULL AND owner_user_id = $1"

	// list by organization
	ProjectListByOrgQuery = "SELECT " +
		"p.id, p.name, p.slug, p.description, p.owner_type, p.owner_user_id, p.owner_org_id, " +
		"p.created_at, p.updated_at " +
		"FROM projects AS p " +
		// filter current user is member of org
		"JOIN rel_org_members AS r ON r.org_id = $1 AND r.user_id = $2 AND r.is_valid = true " +
		"WHERE p.is_valid = true AND p.deleted_at IS NULL AND p.owner_org_id = $1"

	// only by user
	ProjectDetailBySlugQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, deleted_at, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE slug = $1 AND owner_user_id = $2 AND deleted_at IS NULL"

	ProjectDetailBySlugAndOrgIdQuery = "SELECT " +
		"p.id, p.name, p.slug, p.description, p.owner_type, p.owner_user_id, " +
		"p.is_valid, p.deleted_at, " +
		"p.created_at, p.updated_at, " +
		"o.slug, o.name " +
		"FROM projects AS p " +
		"JOIN rel_org_members AS r ON r.user_id = $1 AND r.org_id = $2 AND r.is_valid = true AND r.deleted_at IS NULL " +
		"JOIN orgs AS o ON o.id = $2 AND o.is_valid = true " +
		"WHERE p.slug = $3 AND p.owner_org_id = $2 AND p.deleted_at IS NULL"

	ProjectDetailByIDQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, deleted_at, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE id = $1"

	ProjectCreateQuery = "INSERT INTO " +
		"projects(name, slug, description, owner_type, owner_user_id, owner_org_id) " +
		"VALUES($1, $2, $3, $4, $5, $6) " +
		"RETURNING slug"

	ProjectDeleteQuery = "UPDATE projects " +
		"SET deleted_at = now(), is_valid = false " +
		"WHERE id = $1"

	// @TODO: add query list of slug
)
