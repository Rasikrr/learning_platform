// nolint: revive, stylecheck, unused
package test_cases

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
)

//go:generate mockgen -destination ../mocks/test_cases/mock.go -package mocks -source=repository.go

type Repository interface {
	GetByTaskID(ctx context.Context, taskID string) ([]*entity.TestCase, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(d *databases.Postgres) Repository {
	return &repository{
		db: d,
	}
}

func (r *repository) GetByTaskID(ctx context.Context, taskID string) ([]*entity.TestCase, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByTaskIDStmt, taskID); err != nil {
		return nil, err
	}
	return mm.convert()
}
