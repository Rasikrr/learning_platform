// nolint: revive, stylecheck
package quizzes_submissions

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Repository interface {
	CheckIsPassed(ctx context.Context, userID, topicID string) (bool, error)
	UpdatePassed(ctx context.Context, userID, courseID, topicID string, passed bool) error
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CheckIsPassed(ctx context.Context, userID, topicID string) (bool, error) {
	var exists bool
	if err := pgxscan.Get(ctx, r.db, &exists, checkPassesStmt, userID, topicID); err != nil {
		return false, err
	}
	return exists, nil
}

func (r *repository) UpdatePassed(ctx context.Context, userID, courseID, topicID string, passed bool) error {
	_, err := r.db.Exec(ctx, upsertStmt, userID, courseID, topicID, passed)
	return err
}
