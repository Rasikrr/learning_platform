-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS videos (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    duration INT, -- Длительность в секундах
    position INT NOT NULL, -- Позиция видео в курсе
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF NOT EXISTS videos;
-- +goose StatementEnd
