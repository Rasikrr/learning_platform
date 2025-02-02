-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quizzes (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     topic_id UUID NOT NULL,
     question TEXT NOT NULL,
     options TEXT[] NOT NULL,
     correct_answers BOOLEAN[] NOT NULL,
     created_at TIMESTAMPTZ DEFAULT now(),
     updated_at TIMESTAMPTZ DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quizzes;
-- +goose StatementEnd
