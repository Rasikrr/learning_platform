package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	entities "github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Repository interface {
	GetCourses(ctx context.Context, limit, offset int) ([]*entities.Course, error)
	GetTopicsByIds(ctx context.Context, ids []uuid.UUID) ([]*entities.Topic, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetCourses(ctx context.Context, limit, offset int) ([]*entities.Course, error) {
	var cc courses
	if err := pgxscan.Select(ctx, r.db, &cc, getCoursesStmt, limit, offset); err != nil {
		return nil, err
	}
	return cc.convert()
}

func (r *repository) GetTopicsByIds(ctx context.Context, ids []uuid.UUID) ([]*entities.Topic, error) {
	var tt topics
	if err := pgxscan.Select(ctx, r.db, &tt, getCoursesTopicsByIdsStmt, pq.Array(ids)); err != nil {
		return nil, err
	}
	return tt.convert()
}
