package qs

const (
	/*********************************
		insert
	*********************************/

	UserInsertQuery = "INSERT INTO " +
		"users(id, email, username, is_email_verified, image_url) " +
		"VALUES($1, $2, $3, $4, $5) " +
		"RETURNING id"

	/*********************************
		Update
	*********************************/

	UserUpdateImageOrIsEmailVerifiedQuery = "UPDATE users " +
		"SET is_email_verified = $1, image_url = $2, updated_at = now() " +
		"WHERE id = $3"

	UserUpdateCliPasswordQuery = "UPDATE users " +
		"SET cli_password = $1, updated_at = now() " +
		"WHERE id = $2"

	UserUpdateValidQuery = "UPDATE users " +
		"SET is_valid = $1, updated_at = now() " +
		"WHERE id = $2"

	/*********************************
		Get
	*********************************/

	UserGetByIDQuery = "SELECT " +
		"u.id, u.email, u.username, u.image_url, u.cli_password, u.is_valid, u.is_email_verified, u.created_at, u.updated_at, " +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status " +
		"FROM users AS u " +
		"LEFT JOIN subscriptions AS s ON s.user_id = $1 AND s.is_valid = true AND s.deleted_at IS NULL " +
		"WHERE u.id = $1"

	UserGetByEmailQuery = "SELECT " +
		"u.id, u.email, u.username, u.image_url, u.cli_password, u.is_valid, u.is_email_verified, u.created_at, u.updated_at, " +
		"s.id, s.stripe_subscription_id, s.stripe_customer_id, " +
		"s.stripe_product_id, s.stripe_subscription_status " +
		"FROM users AS u " +
		"LEFT JOIN subscriptions AS s ON s.user_id = u.id AND s.is_valid = true AND s.deleted_at IS NULL " +
		"WHERE u.email = $1"

	UserExistsQuery = "SELECT " +
		"id, is_valid, is_email_verified, image_url " +
		"FROM users " +
		"WHERE id = $1"

	UserGetByEmailAdnPassword = "SELECT " +
		"id, email, username, image_url, cli_password, created_at, updated_at " +
		"FROM users " +
		"WHERE (email = $1 OR username = $1) AND is_valid = true AND is_email_verified"
)
