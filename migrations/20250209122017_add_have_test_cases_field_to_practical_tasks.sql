-- +goose Up
-- +goose StatementBegin
ALTER TABLE practical_tasks
    ADD COLUMN test_cases BOOLEAN NOT NULL DEFAULT false;

ALTER TABLE practical_tasks
    ADD COLUMN language VARCHAR(255) NOT NULL DEFAULT 'go';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE practical_tasks
DROP COLUMN test_cases;

ALTER TABLE practical_tasks
DROP COLUMN language;
-- +goose StatementEnd
