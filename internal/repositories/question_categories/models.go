package question_categories

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type models []model

func (m model) convert() (*entity.QuestionCategory, error) {
	return &entity.QuestionCategory{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func (m models) convert() ([]*entity.QuestionCategory, error) {
	var err error
	out := make([]*entity.QuestionCategory, len(m))
	for i, item := range m {
		out[i], err = item.convert()
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
