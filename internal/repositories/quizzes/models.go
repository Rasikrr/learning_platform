package quizzes

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID             uuid.UUID
	TopicID        uuid.UUID
	Question       string
	Options        []string
	CorrectAnswers []bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type models []model

// nolint: unparam
func (m model) convert() (*entity.Quiz, error) {
	return &entity.Quiz{
		ID:             m.ID,
		TopicID:        m.TopicID,
		Question:       m.Question,
		Options:        m.Options,
		CorrectAnswers: m.CorrectAnswers,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.Quiz, error) {
	out := make([]*entity.Quiz, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
