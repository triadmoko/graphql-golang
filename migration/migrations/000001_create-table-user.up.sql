CREATE TABLE users (
    id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    email VARCHAR NOT NULL,
    name VARCHAR NULL,
    password VARCHAR NOT NULL,
    status VARCHAR NOT NULL DEFAULT 'inactive',
    PRIMARY KEY (id)
);