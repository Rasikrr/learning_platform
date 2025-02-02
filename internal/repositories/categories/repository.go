package categories

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*entity.Category, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]*entity.Category, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]*entity.Category, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getCategoriesByIdsStmt, pq.Array(ids)); err != nil {
		return nil, err
	}
	return mm.convert()
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.Category, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getAllCategoriesStmt); err != nil {
		return nil, err
	}
	return mm.convert()
}
