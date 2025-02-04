-- +goose Up
-- +goose StatementBegin
CREATE TABLE faq_questions (
    id uuid,
    category_id uuid NOT NULL,
    user_id uuid NOT NULL,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT faq_questions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS faq_categories (
    id uuid,
    name text NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT faq_categories_pkey PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS faq_answers (
    id uuid,
    question_id uuid NOT NULL,
    user_id uuid NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT faq_answers_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS faq_answers;
DROP TABLE IF EXISTS faq_questions;
DROP TABLE IF EXISTS faq_categories;
-- +goose StatementEnd
