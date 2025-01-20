-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN account_role VARCHAR(255) NOT NULL DEFAULT 'user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN account_role;
-- +goose StatementEnd
