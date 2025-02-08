package answers

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, answer *entity.Answer) error
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
