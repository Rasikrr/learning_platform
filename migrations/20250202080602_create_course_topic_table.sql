-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS course_topics (
    id uuid,
    name VARCHAR(255) NOT NULL,
    created_by uuid NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT course_topics_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS course_topics;
-- +goose StatementEnd
