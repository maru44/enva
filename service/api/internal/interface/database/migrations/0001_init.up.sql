BEGIN;

-- organization
CREATE TABLE orgs (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    slug VARCHAR(63) NOT NULL,
    name VARCHAR(63) NOT NULL,
    description VARCHAR(511) NULL,
    is_valid BOOLEAN NOT NULL,
    created_by VARCHAR(127) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);

-- @TODO add username and email and cli password table
-- OR make custome field to aws cognito

-- CREATE TABLE users (
--     id VARCHAR(127),
--     email VARCHAR(255),
--     username VARCHAR(31),
--     image_url VARCHAR(255),
--     is_valid BOOLEAN NOT NULL,
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

--     PRIMARY KEY (id)
-- );

-- cli password

-- if user is deleted >> is_valid must be changed to false
CREATE TABLE cli_users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    cognito_id VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(63) NOT NULL,
    password VARCHAR(512),
    is_valid BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);
CREATE INDEX ON cli_users (email);
CREATE INDEX ON cli_users (username);

-- relation org and users
CREATE TABLE rel_org_members (
    org_id uuid NOT NULL REFERENCES orgs (id) ON DELETE CASCADE,
    user_id VARCHAR(127) NOT NULL,
    user_type VARCHAR(15) NOT NULL DEFAULT 'user',
    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    PRIMARY KEY (org_id, user_id)
);

-- projects
CREATE TABLE projects (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(63) NOT NULL,
    slug VARCHAR(63) NOT NULL,
    description TEXT NULL,
    owner_type VARCHAR(15) NOT NULL,
    owner_user_id VARCHAR(127) NULL, -- dbのユーザーのidのfkにする?
    owner_org_id uuid NULL REFERENCES orgs (id) ON DELETE CASCADE,
    is_valid BOOLEAN NOT NULL DEFAULT true,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id),
    CONSTRAINT slug_unique_every_user UNIQUE (slug, owner_user_id) WHERE (is_valid = true),
    CONSTRAINT slug_unique_every_org UNIQUE (slug, owner_org_id) WHERE (is_valid = true)
);
CREATE INDEX ON projects (owner_user_id);

-- key value sets
CREATE TABLE kvs (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    project_id uuid NOT NULL REFERENCES projects (id) ON DELETE CASCADE,
    env_key VARCHAR(255) NOT NULL,
    env_value VARCHAR(1023) NOT NULL DEFAULT '',
    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    created_by VARCHAR(127) NOT NULL,
    updated_by VARCHAR(127) NULL DEFAULT NULL,

    PRIMARY KEY (id)
);
CREATE INDEX ON kvs (env_key, project_id, is_valid);
CREATE INDEX ON kvs (project_id, is_valid);
CREATE UNIQUE INDEX valid_project_kvs ON kvs (env_key, project_id) WHERE (is_valid = true);

COMMIT;
