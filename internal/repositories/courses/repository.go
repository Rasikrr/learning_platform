package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/georgysavva/scany/v2/pgxscan"
	"log"
)

type Repository interface {
	GetCourses(ctx context.Context, limit, offset int) error
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetCourses(ctx context.Context, limit, offset int) error {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getCoursesStmt, limit, offset); err != nil {
		return err
	}
	for _, m := range mm {
		log.Println(m)
	}
	return nil
}
