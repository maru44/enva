package qs

const (
	/*********************************
		List
	*********************************/

	ValidKvListOfProject = "SELECT " +
		"id, env_key, env_value, is_valid, created_at, updated_at " +
		"FROM kvs " +
		"WHERE project_id = $1 AND is_valid = true"

	/*********************************
		Detail
	*********************************/

	ValidKvDetail = "SELECT " +
		"id, env_key, env_value, is_valid, created_at, updated_at " +
		"FROM kvs " +
		"WHERE env_key = $1 AND project_id = $2 AND is_valid = true"

	ValidKvDetailID = "SELECT " +
		"id " +
		"FROM kvs " +
		"WHERE env_key = $1 AND project_id = $2 AND is_valid = true"

	/*********************************
		Insert
	*********************************/

	KvInsertQuery = "INSERT INTO " +
		"kvs(env_key, env_value, project_id, created_by) " +
		"VALUES($1, $2, $3, $4) " +
		"RETURNING id"

	/*********************************
		Delete
	*********************************/

	KvDeactivateQuery = "UPDATE kvs " +
		"SET is_valid = false, updated_by = $1, updated_at = now() " +
		"WHERE project_id = $2 AND env_key = $3 AND is_valid = true"

	KvDeactivateByIdQuery = "UPDATE kvs " +
		"SET is_valid = false, updated_by = $1, updated_at = now() " +
		"WHERE project_id = $2 AND id = $3 AND is_valid = true"
)
