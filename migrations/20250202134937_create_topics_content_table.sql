-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS topic_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    topic_id UUID NOT NULL,
    content TEXT NOT NULL,
    additional_resources TEXT[],
    video_urls TEXT[],
    image_urls TEXT[],
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);




-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS topic_contents;
DROP TABLE IF EXISTS topic_videos;
-- +goose StatementEnd
