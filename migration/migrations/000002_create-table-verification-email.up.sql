CREATE TABLE email_verifications (
    id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    user_id uuid NOT NULL,
    email VARCHAR NOT NULL,
    code INT NOT NULL,
    expired BIGINT NOT NULL,
    PRIMARY KEY (id)
);