package content

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/lib/pq"
)

type Repository interface {
	GetByTopicID(ctx context.Context, id string) (*entity.TopicContent, error)
	GetByTopicIDs(ctx context.Context, topicIDs []string) ([]*entity.TopicContent, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByTopicIDs(ctx context.Context, topicIDs []string) ([]*entity.TopicContent, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByTopicIDsStmt, pq.Array(topicIDs)); err != nil {
		return nil, err
	}
	return mm.convert()
}

func (r *repository) GetByTopicID(ctx context.Context, id string) (*entity.TopicContent, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getByIDStmt, id); err != nil {
		return nil, err
	}
	return m.convert()
}
