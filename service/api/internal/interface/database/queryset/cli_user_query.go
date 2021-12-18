package queryset

const (
	CliUserInsertQuery = "INSET INTO " +
		"cli_users(email, username, password) " +
		"VALUES($1, $2, $3) " +
		"RETURNING id"

	CliUserUpdateQuery = "UPDATE cli_users " +
		"SET password = $1, is_valid = true " +
		"WHERE email = $1"

	CliUserDeactivateQuery = "UPDATE cli_users " +
		"SET is_valid = false " +
		"WHERE email = $1 AND is_valid = true"

	CliUserExistsQuery = "SELECT id " +
		"FROM cli_users " +
		"WHERE email = $1 AND is_valid = true"

	CliUserGetPasswordByEmailOrPassword = "SELECT password " +
		"FROM cli_users " +
		"WHERE email = $1 AND username = $1 AND is_valid = true"
)
