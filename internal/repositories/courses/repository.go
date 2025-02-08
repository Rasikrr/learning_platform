package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Repository interface {
	GetByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error)
	GetByID(ctx context.Context, id string) (*entity.Course, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error) {
	query, args, err := generateQuery(params)
	if err != nil {
		return nil, err
	}
	var cc models
	if err = pgxscan.Select(ctx, r.db, &cc, query, args...); err != nil {
		return nil, err
	}
	return cc.convert()
}

func (r *repository) GetByID(ctx context.Context, id string) (*entity.Course, error) {
	var c model
	if err := pgxscan.Get(ctx, r.db, &c, getCourseByIDStmt, id); err != nil {
		return nil, err
	}
	return c.convert()
}
