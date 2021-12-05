CREATE TABLE public.posts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    title CHARACTER VARYING(63) NOT NULL,
    content CHARACTER VARYING(511) NULL,
    abstract CHARACTER VARYING(511) NOT NULL,
    is_public BOOLEAN NOT NULL,
    is_valid BOOLEAN NOT NULL,
    user_id CHARACTER VARYING(63) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE public.users (
    id CHARACTER VARYING(63),
    email CHARACTER VARYING(255),
    username CHARACTER VARYING(31),
    image_url CHARACTER VARYING(255),
    is_valid BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
