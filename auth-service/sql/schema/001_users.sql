-- +goose Up


CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name character varying(255),
    last_name character varying(255),
    email_address TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    user_active INTEGER DEFAULT 0,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
    );

-- +goose Down

DROP TABLE IF EXISTS users;