package answers

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID         uuid.UUID
	QuestionID uuid.UUID
	UserID     uuid.UUID
	Body       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type models []model

func (m model) convert() (*entity.Answer, error) {
	question := entity.Question{
		ID: m.QuestionID,
	}
	user := entity.User{
		ID: m.UserID,
	}
	return &entity.Answer{
		ID:        m.ID,
		Author:    &user,
		Question:  &question,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func (m models) convert() ([]*entity.Answer, error) {
	var err error
	out := make([]*entity.Answer, len(m))
	for i, item := range m {
		out[i], err = item.convert()
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
