-- +goose Up
-- +goose StatementBegin
ALTER TABLE course_enrollments
    ADD CONSTRAINT course_enrollments_course_id_user_id_unique UNIQUE (course_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE course_enrollments
    DROP CONSTRAINT course_enrollments_course_id_user_id_unique;
-- +goose StatementEnd
