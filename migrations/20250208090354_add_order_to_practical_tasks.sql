-- +goose Up
-- +goose StatementBegin
ALTER TABLE practical_tasks
    ADD COLUMN order_number INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE practical_tasks
    DROP COLUMN order_number;
-- +goose StatementEnd
