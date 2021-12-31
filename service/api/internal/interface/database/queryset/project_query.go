package queryset

const (
	ProjectListQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at = null"

	ProjectListByUserQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at = null AND owner_user_id = $1"

	ProjectSlugListByUserQuery = "SELECT " +
		"slug " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at = null AND owner_user_id = $1"

	// list by organization
	ProjectListByOrgQuery = "SELECT " +
		"id, name, slug, description, owenr_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND deleted_at = null AND owner_org_id = $1"

	// only by user
	ProjectDetailBySlugQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, deleted_at, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE slug = $1 AND owner_user_id = $2 AND deleted_at = null"

	ProjectDetailBySlugAndOrgIdQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, deleted_at, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE slug = $1 AND owner_org_id = $2 AND deleted_at = null"

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
