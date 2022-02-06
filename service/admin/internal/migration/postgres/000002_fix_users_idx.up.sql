BEGIN;

DROP INDEX users_email_idx;
DROP INDEX users_username_idx;

CREATE UNIQUE INDEX users_email_valid_idx ON users (email) WHERE (is_valid = true);
CREATE UNIQUE INDEX users_username_valid_idx ON users (username) WHERE (is_valid = true);

COMMIT;
