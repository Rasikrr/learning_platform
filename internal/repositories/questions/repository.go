package question

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/databases"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
)

type Repository interface {
	Create(ctx context.Context, question *entity.Question) error
}

type repository struct {
	db *databases.Postgres
}

func NewRepository(db *databases.Postgres) Repository {
	return &repository{db}
}

func (r *repository) Create(ctx context.Context, question *entity.Question) error {
	_, err := r.db.Exec(ctx, createQuestionStmt, question.ID, question.Category.ID, question.Author.ID, question.Title, question.Body)
	if err != nil {
		return err
	}
	return nil
}
