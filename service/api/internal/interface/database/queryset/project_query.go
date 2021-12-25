package queryset

const (
	ProjectListQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND is_deleted = false"

	ProjectListByUserQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND is_deleted = false AND owner_user_id = $1"

	ProjectSlugListByUserQuery = "SELECT " +
		"slug " +
		"FROM projects " +
		"WHERE is_valid = true AND is_deleted = false AND owner_user_id = $1"

	ProjectListByOrgQuery = "SELECT " +
		"id, name, slug, description, owenr_type, owner_org_id, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE is_valid = true AND is_deleted = false AND owner_org_id = $1"

	ProjectDetailBySlugQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, is_deleted, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE slug = $1 AND is_deleted = false"

	ProjectDetailByIDQuery = "SELECT " +
		"id, name, slug, description, owner_type, owner_user_id, owner_org_id, " +
		"is_valid, is_deleted, " +
		"created_at, updated_at " +
		"FROM projects " +
		"WHERE id = $1"

	ProjectCreateQuery = "INSERT INTO " +
		"projects(name, slug, description, owner_type, owner_user_id, owner_org_id) " +
		"VALUES($1, $2, $3, $4, $5, $6) " +
		"RETURNING slug"

	ProjectDeleteQuery = "UPDATE projects " +
		"SET is_deleted = true, is_valid = false " +
		"WHERE id = $1"

	// @TODO: add query list of slug
)
