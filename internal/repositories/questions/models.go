// nolint: unused
package question

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID         uuid.UUID
	CategoryID uuid.UUID
	UserID     uuid.UUID
	ImageURL   *string
	Title      string
	Body       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type models []model

func (m model) convert() (*entity.Question, error) {
	category := entity.QuestionCategory{
		ID: m.CategoryID,
	}
	user := entity.User{
		ID: m.UserID,
	}
	return &entity.Question{
		ID:        m.ID,
		Category:  &category,
		Author:    &user,
		ImageURL:  m.ImageURL,
		Title:     m.Title,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func (m models) convert() ([]*entity.Question, error) {
	var err error
	out := make([]*entity.Question, len(m))
	for i, item := range m {
		out[i], err = item.convert()
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
