-- +goose Up
-- +goose StatementBegin
ALTER TABLE faq_questions
    ADD COLUMN image_url TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE faq_questions
    DROP COLUMN image_url;
-- +goose StatementEnd
