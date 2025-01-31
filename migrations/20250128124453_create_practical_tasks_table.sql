-- +goose Up
-- +goose StatementBegin
CREATE TABLE practical_tasks (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    sample_code TEXT, -- Пример кода
    solution_code TEXT, -- Решение
    position INT NOT NULL, -- Позиция задания в курсе
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS practical_tasks;
-- +goose StatementEnd
