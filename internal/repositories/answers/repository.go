package answers

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Repository interface {
	Create(ctx context.Context, answer *entity.Answer) error
	GetByQuestionID(ctx context.Context, questionID string, limit, offset int) ([]*entity.Answer, error)
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{db}
}

func (r *repository) Create(ctx context.Context, answer *entity.Answer) error {
	_, err := r.db.Exec(ctx, createAnswerStmt, answer.ID, answer.Question.ID, answer.Author.ID, answer.Body)
	return err
}

func (r *repository) GetByQuestionID(ctx context.Context, questionID string, limit, offset int) ([]*entity.Answer, error) {
	var mm models
	if err := pgxscan.Select(ctx, r.db, &mm, getByQuestionIDStmt, questionID, limit, offset); err != nil {
		return nil, err
	}
	return mm.convert()
}
