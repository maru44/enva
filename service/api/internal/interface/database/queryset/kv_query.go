package queryset

const (
	ValidKvListOfProject = "SELECT " +
		"id, env_key, env_value, is_valid, created_at, updated_at " +
		"FROM kvs " +
		"WHERE project_id = $1 AND is_valid = true"

	ValidKvDetail = "SELECT " +
		"id, env_key, env_value, is_valid, created_at, updated_at " +
		"FROM kvs " +
		"WHERE env_key = $1 AND project_id = $2 AND is_valid = true"

	KvInsertQuery = "INSERT INTO " +
		"kvs(env_key, env_value, project_id, created_by) " +
		"VALUES($1, $2, $3, $4) " +
		"RETURNING id"

	ValidKvDetailID = "SELECT " +
		"id " +
		"FROM kvs " +
		"WHERE env_key = $1 AND project_id = $2 AND is_valid = true"

	KvDeactivateQuery = "UPDATE kvs " +
		"SET is_valid = false, updated_by = $1 " +
		"WHERE project_id = $2 AND env_key = $3 AND is_valid = true"

	KvDeactivateByIdQuery = "UPDATE kvs " +
		"SET is_valid = false, updated_by = $1 " +
		"WHERE project_id = $2 AND id = $3 AND is_valid = true"
)
