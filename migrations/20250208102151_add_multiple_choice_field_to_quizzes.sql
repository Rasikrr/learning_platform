-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes
    ADD COLUMN multiple_choice BOOLEAN NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE quizzes
    DROP COLUMN multiple_choice;
-- +goose StatementEnd
