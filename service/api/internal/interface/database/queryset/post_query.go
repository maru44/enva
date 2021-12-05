package queryset

const (
	// get post list
	PostListQuery = "SELECT id, title, abstract, content, is_valid, is_public, created_at, updated_at, user_id " +
		"FROM posts " +
		"ORDER BY created_at DESC"

	PostInsertQuery = "INSERT INTO " +
		"posts(title, abstract, content, is_public, is_valid, user_id) " +
		"VALUES($1, $2, $3, $4, $5, $6) " +
		"RETURNING id"
)
