-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS practical_tasks (
     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
     topic_id UUID NOT NULL,
     description TEXT NOT NULL,
     difficulty_level VARCHAR(255) NOT NULL,
     starter_code TEXT,
     expected_output TEXT,
     created_at TIMESTAMPTZ DEFAULT now(),
     updated_at TIMESTAMPTZ DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS practical_tasks;
-- +goose StatementEnd
