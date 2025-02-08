-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quizzes_submissions (
    id uuid DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    course_id uuid NOT NULL,
    topic_id uuid NOT NULL,
    passed boolean NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT quiz_submissions_pkey PRIMARY KEY (id),
    CONSTRAINT quiz_submissions_user_topic UNIQUE (user_id, topic_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quizes_submissions;
-- +goose StatementEnd
