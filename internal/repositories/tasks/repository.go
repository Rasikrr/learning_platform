package tasks

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/lib/pq"
)

//go:generate mockgen -destination ../mocks/tasks/mock.go -package mocks -source=repository.go

type Repository interface {
	GetByID(ctx context.Context, id string) (*entity.PracticalTask, error)
	GetByTopicIDAndOrderNum(ctx context.Context, id string, order int) (*entity.PracticalTask, error)
	GetByTopicIDs(ctx context.Context, ids []string) ([]*entity.PracticalTask, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByID(ctx context.Context, id string) (*entity.PracticalTask, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getByIDStmt, id); err != nil {
		return nil, err
	}
	return m.convert()
}

func (r *repository) GetByTopicIDAndOrderNum(ctx context.Context, id string, order int) (*entity.PracticalTask, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getByTopicIDAndOrderNumStmt, id, order); err != nil {
		return nil, err
	}
	return m.convert()
}

func (r *repository) GetByTopicIDs(ctx context.Context, ids []string) ([]*entity.PracticalTask, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByTopicIDsStmt, pq.Array(ids)); err != nil {
		return nil, err
	}
	return mm.convert()
}
