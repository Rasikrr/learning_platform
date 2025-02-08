package tasks

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID              uuid.UUID
	TopicID         uuid.UUID
	Description     string
	DifficultyLevel string
	StarterCode     string
	ExpectedOutput  string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	OrderNumber     int
}

type models []model

func (m model) convert() (*entity.PracticalTask, error) {
	level, err := enum.DifficultyString(m.DifficultyLevel)
	if err != nil {
		return nil, err
	}
	return &entity.PracticalTask{
		ID:              m.ID,
		TopicID:         m.TopicID,
		Description:     m.Description,
		DifficultyLevel: level,
		StarterCode:     m.StarterCode,
		ExpectedOutput:  m.ExpectedOutput,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		OrderNumber:     m.OrderNumber,
	}, nil
}

func (mm models) convert() ([]*entity.PracticalTask, error) {
	out := make([]*entity.PracticalTask, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
