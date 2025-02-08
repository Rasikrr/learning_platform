-- +goose Up
-- +goose StatementBegin
ALTER TABLE topics
    ADD COLUMN description TEXT NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE topics
    DROP COLUMN description;
-- +goose StatementEnd
