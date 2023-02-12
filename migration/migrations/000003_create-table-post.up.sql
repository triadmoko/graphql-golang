CREATE TABLE posts (
    id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    user_id uuid NOT NULL,
    title VARCHAR NULL,
    description VARCHAR NOT NULL,
    PRIMARY KEY (id)
);