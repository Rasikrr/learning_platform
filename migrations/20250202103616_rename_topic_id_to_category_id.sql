-- +goose Up
-- +goose StatementBegin
ALTER TABLE courses
    RENAME COLUMN topic_id TO category_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE courses
    RENAME COLUMN category_id TO topic_id;
-- +goose StatementEnd
