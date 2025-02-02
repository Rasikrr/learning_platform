-- +goose Up
-- +goose StatementBegin
ALTER TABLE course_topics
    RENAME TO course_category;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE course_category
    RENAME TO course_topics;
-- +goose StatementEnd
