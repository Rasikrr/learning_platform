// nolint: revive, stylecheck, unused
package tasks_submissions

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
)

//go:generate mockgen -destination ../mocks/tasks_submissions/mock.go -package mocks -source=repository.go

type Repository interface {
	Create(ctx context.Context, submission *entity.TaskSubmission) error
	CheckIsPassed(ctx context.Context, userID, taskID string) (bool, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, submission *entity.TaskSubmission) error {
	m := convert(submission)
	_, err := r.db.Exec(ctx, insertStmt, m.TaskID, m.UserID, m.Input, m.Passed, m.Error)
	return err
}

func (r *repository) CheckIsPassed(ctx context.Context, userID, taskID string) (bool, error) {
	var exists bool
	if err := pgxscan.Get(ctx, r.db, &exists, checkPassesStmt, userID, taskID); err != nil {
		return false, err
	}
	return exists, nil
}
