package topics

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

type Repository interface {
	GetByCourseID(ctx context.Context, id uuid.UUID) ([]*entity.Topic, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByCourseID(ctx context.Context, ids uuid.UUID) ([]*entity.Topic, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByCourseIDStmt, ids); err != nil {
		return nil, err
	}
	return mm.convert()
}
