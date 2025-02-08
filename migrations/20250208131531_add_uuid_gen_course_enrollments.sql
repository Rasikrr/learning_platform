-- +goose Up
-- +goose StatementBegin
ALTER TABLE course_enrollments
    ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE course_enrollments
    ALTER COLUMN id DROP DEFAULT;
-- +goose StatementEnd
