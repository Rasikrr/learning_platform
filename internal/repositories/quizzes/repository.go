package quizzes

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/lib/pq"
)

//go:generate mockgen -destination ../mocks/quizzes/mock.go -package mocks -source=repository.go

type Repository interface {
	GetByTopicID(ctx context.Context, id string) ([]*entity.Quiz, error)
	GetByTopicIDs(ctx context.Context, ids []string) ([]*entity.Quiz, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByTopicID(ctx context.Context, id string) ([]*entity.Quiz, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByTopicIDStmt, id); err != nil {
		return nil, err
	}
	return mm.convert()
}

func (r *repository) GetByTopicIDs(ctx context.Context, ids []string) ([]*entity.Quiz, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByTopicIDsStmt, pq.Array(ids)); err != nil {
		return nil, err
	}
	return mm.convert()
}
