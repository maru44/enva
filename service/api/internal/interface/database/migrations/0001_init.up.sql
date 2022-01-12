BEGIN;

-- @TODO add username and email and cli password table
-- OR make custome field to aws cognito

CREATE TABLE users (
    id uuid NOT NULL,
    email VARCHAR(255),
    username VARCHAR(31),
    image_url VARCHAR(255),
    cli_password VARCHAR(511) DEFAULT NULL,
    is_valid BOOLEAN NOT NULL DEFAULT true,
    is_email_verified BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX ON users (email);
CREATE UNIQUE INDEX ON users (username);

-- organization
CREATE TABLE orgs (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    slug VARCHAR(63) UNIQUE NOT NULL,
    name VARCHAR(63) NOT NULL,
    description VARCHAR(511) NULL,
    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_by uuid NULL REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE subscriptions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NULL REFERENCES users (id) ON DELETE CASCADE,
    org_id uuid NULL REFERENCES orgs (id) ON DELETE CASCADE,

    stripe_subscription_id VARCHAR(63) NOT NULL,
    stripe_customer_id VARCHAR(63) NOT NULL,
    stripe_product_id VARCHAR(63) NOT NULL,
    stripe_subscription_status VARCHAR(63) DEFAULT NULL,

    is_valid BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    PRIMARY KEY (id),
    CONSTRAINT subscription_owner CHECK (user_id IS NOT NULL OR org_id IS NOT NULL)
);

CREATE TABLE ssh_keys (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(31) NOT NULL,
    pub_key TEXT NOT NULL,
    user_id uuid REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE org_invitations (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    org_id uuid NOT NULL REFERENCES orgs (id) ON DELETE CASCADE,
    user_id uuid REFERENCES users (id) DEFAULT NULL,
    email VARCHAR(255) NOT NULL,
    user_type VARCHAR(15) NOT NULL DEFAULT 'user', -- 'owner', 'admin', 'user', 'guest'
    status VARCHAR(15) NOT NULL DEFAULT 'new', -- 'accepted', 'denied', 'new', 'closed'
    invited_by uuid REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);

-- relation org and users
CREATE TABLE rel_org_members (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    org_id uuid NOT NULL REFERENCES orgs (id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users (id),
    user_type VARCHAR(15) NOT NULL DEFAULT 'user', -- 'owner', 'admin', 'user', 'guest'

    org_invitation_id uuid REFERENCES org_invitations (id) DEFAULT NULL,

    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX rel_org_members_user_id_org_id ON rel_org_members (user_id, org_id) WHERE (is_valid = true);

-- projects
CREATE TABLE projects (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name VARCHAR(63) NOT NULL,
    slug VARCHAR(63) NOT NULL,
    description TEXT NULL,
    owner_type VARCHAR(15) NOT NULL,
    owner_user_id uuid NULL REFERENCES users (id),
    owner_org_id uuid NULL REFERENCES orgs (id) ON DELETE CASCADE,
    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    PRIMARY KEY (id),
    CONSTRAINT projeccts_owner CHECK (owner_user_id IS NOT NULL OR owner_org_id IS NOT NULL)
);
CREATE UNIQUE INDEX project_slug_owner_user ON projects (slug, owner_user_id) WHERE (is_valid = true);
CREATE UNIQUE INDEX project_slug_owner_org_id ON projects (slug, owner_org_id) WHERE (is_valid = true);

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
-- CREATE INDEX ON kvs (env_key, project_id, is_valid);
CREATE INDEX ON kvs (project_id, is_valid);
CREATE UNIQUE INDEX kvs_env_key_project_id_valid ON kvs (env_key, project_id) WHERE (is_valid = true);

COMMIT;
