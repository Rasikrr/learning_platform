// nolint: revive, stylecheck, unused
package tasks_submissions

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

//go:generate mockgen -destination ../mocks/tasks_submissions/mock.go -package mocks -source=repository.go

type Repository interface {
	Create(ctx context.Context, submission *entity.TaskSubmission) error
	CheckIsPassed(ctx context.Context, userID, taskID string) (bool, error)
	DeleteByUserAndTaskID(ctx context.Context, userID, taskID string) error
	GetSolvedTasks(ctx context.Context, userID, taskID string) (*entity.TaskSubmission, bool, error)
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

func (r *repository) DeleteByUserAndTaskID(ctx context.Context, userID, taskID string) error {
	_, err := r.db.Exec(ctx, deleteByUserAndTaskIDStmt, userID, taskID)
	return err
}

func (r *repository) GetSolvedTasks(ctx context.Context, userID, taskID string) (*entity.TaskSubmission, bool, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getSolvedTasksStmt, userID, taskID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}
	e, err := m.convert()
	if err != nil {
		return nil, false, err
	}
	return e, true, nil
}
