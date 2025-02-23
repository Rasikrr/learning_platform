package question

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Create(ctx context.Context, question *entity.Question) error
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Question, error)
	GetByParams(ctx context.Context, params *entity.GetQuestionsParams) ([]*entity.Question, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{db}
}

func (r *repository) Create(ctx context.Context, question *entity.Question) error {
	_, err := r.db.Exec(ctx, createQuestionStmt, question.ID, question.Category.ID,
		question.Author.ID, question.Title, question.Body, question.ImageURL)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Question, error) {
	var m model
	if err := pgxscan.Get(ctx, r.db, &m, getByIDStmt, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("question not found")
		}
		return nil, err
	}
	return m.convert()
}

func (r *repository) GetByParams(ctx context.Context, params *entity.GetQuestionsParams) ([]*entity.Question, error) {
	query, args, err := generateQuery(params)
	if err != nil {
		return nil, err
	}
	var mm models
	if err = pgxscan.Select(ctx, r.db, &mm, query, args...); err != nil {
		return nil, err
	}
	return mm.convert()
}
