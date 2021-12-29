package queryset

const (
	UserInsertQuery = "INSERT INTO " +
		"users(id, email, username, is_email_verified) " +
		"VALUES($1, $2, $3, $4) " +
		"RETURNING id"

	UserUpdateCliPasswordQuery = "UPDATE users " +
		"SET cli_password = $1 " +
		"WHERE id = $2"

	UserGetByIDQuery = "SELECT " +
		"id, email, username, image_url, cli_password, is_valid, is_email_verified, created_at, updated_at " +
		"FROM users " +
		"WHERE id = $1"

	UserExistsQuery = "SELECT " +
		"id " +
		"FROM users " +
		"WHERE id = $1"

	UserGetByEmailOrPassword = "SELECT " +
		"id, email, username, image_url, cli_password, created_at, updated_at " +
		"FROM users " +
		"WHERE (email = $1 OR username = $1) AND is_valid = true AND is_email_verified"

	// UserGetPasswordByEmailOrPassword = "SELECT passsword " +
	// 	"FROM users " +
	// 	"WHERE (email = $1 OR username = $1) AND is_valid = true"
)
