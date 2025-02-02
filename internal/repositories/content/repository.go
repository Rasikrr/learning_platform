package content

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository interface {
	GetByTopicIDs(ctx context.Context, topicIDs []uuid.UUID) ([]*entity.TopicContent, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByTopicIDs(ctx context.Context, topicIDs []uuid.UUID) ([]*entity.TopicContent, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByTopicIDsStmt, pq.Array(topicIDs)); err != nil {
		return nil, err
	}
	return mm.convert()
}
