-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS practical_tasks_submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    task_id UUID NOT NULL,
    user_id UUID NOT NULL,
    input TEXT NOT NULL,
    passed BOOLEAN NOT NULL DEFAULT false,
    error TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS practical_tasks_submissions;
-- +goose StatementEnd
