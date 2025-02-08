-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS course_enrollments (
    id uuid,
    user_id uuid NOT NULL,
    course_id uuid NOT NULL,
    status VARCHAR(255) NOT NULL DEFAULT 'in_progress',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT course_enrollments_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS course_enrollments;
-- +goose StatementEnd
