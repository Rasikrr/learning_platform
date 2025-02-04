// nolint: revive, stylecheck
package question_categories

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.QuestionCategory, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{db}
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*entity.QuestionCategory, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getByIDStmt, id); err != nil {
		return nil, err
	}

	return m.convert()
}
