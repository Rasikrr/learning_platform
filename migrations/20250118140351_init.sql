-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id uuid,
    name text,
    last_name text,
    email text NOT NULL UNIQUE,
    password text,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    deleted_at timestamptz,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
