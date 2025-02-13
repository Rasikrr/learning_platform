package enrollments

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Repository interface {
	Enroll(ctx context.Context, userID, courseID string) error
	GetUserEnrollments(cxt context.Context, userID string) ([]*entity.Enrollment, error)
	CheckByUserIDAndCourseID(ctx context.Context, userID, courseID string) (bool, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Enroll(ctx context.Context, userID, courseID string) error {
	_, err := r.db.Exec(ctx, enrollUserStmt, userID, courseID)
	return err
}

func (r *repository) CheckByUserIDAndCourseID(ctx context.Context, userID, courseID string) (bool, error) {
	var exists bool
	if err := pgxscan.Get(ctx, r.db, &exists, getByUserIDAndCourseIDStmt, userID, courseID); err != nil {
		return false, err
	}
	return exists, nil
}

func (r *repository) GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getUserEnrollmentsStmt, userID); err != nil {
		return nil, err
	}
	return mm.convert()
}
