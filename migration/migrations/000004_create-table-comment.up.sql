CREATE TABLE comments (
    id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    user_id uuid NOT NULL,
    post_id uuid NOT NULL,
    description VARCHAR NOT NULL,
    PRIMARY KEY (id)
);