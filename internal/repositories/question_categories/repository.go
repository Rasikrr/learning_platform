// nolint: revive, stylecheck
package question_categories

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*entity.QuestionCategory, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.QuestionCategory, error)
	GetByIDs(ctx context.Context, ids []string) ([]*entity.QuestionCategory, error)
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return m.convert()
}

func (r *repository) GetAll(ctx context.Context) ([]*entity.QuestionCategory, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getAllStmt); err != nil {
		return nil, err
	}
	return mm.convert()
}

func (r *repository) GetByIDs(ctx context.Context, ids []string) ([]*entity.QuestionCategory, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByIDsStmt, pq.Array(ids)); err != nil {
		return nil, err
	}
	return mm.convert()
}
