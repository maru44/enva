BEGIN;

DROP INDEX users_email_valid_idx;
DROP INDEX users_username_valid_idx;

CREATE UNIQUE INDEX ON users (email);
CREATE UNIQUE INDEX ON users (username);

COMMIT;
