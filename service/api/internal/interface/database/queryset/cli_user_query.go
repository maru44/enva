package queryset

const (
	CliUserInsertQuery = "INSERT INTO " +
		"cli_users(cognito_id, email, username, password) " +
		"VALUES($1, $2, $3, $4) " +
		"RETURNING id"

	CliUserUpdateQuery = "UPDATE cli_users " +
		"SET password = $1, is_valid = true " +
		"WHERE email = $2"

	CliUserDeactivateQuery = "UPDATE cli_users " +
		"SET is_valid = false " +
		"WHERE email = $1 AND is_valid = true"

	CliUserExistsQuery = "SELECT id " +
		"FROM cli_users " +
		"WHERE email = $1 AND is_valid = true"

	CliUserGetPasswordByEmailOrPassword = "SELECT password " +
		"FROM cli_users " +
		// "WHERE (email = $1 OR username = $1) AND is_valid = true"
		"WHERE email = $1 AND is_valid = true"

	CliUserGet = "SELECT cognito_id, email, username, password " +
		"FROM cli_users " +
		// "WHERE (email = $1 OR username = $1) AND is_valid = true"
		"WHERE email = $1 AND is_valid = true"
)
