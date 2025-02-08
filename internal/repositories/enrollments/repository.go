package enrollments

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Repository interface {
	Enroll(ctx context.Context, userID, courseID string) error
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
