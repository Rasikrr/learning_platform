package topics

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID          uuid.UUID
	CourseID    uuid.UUID
	Title       string
	OrderNumber int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type models []model

// nolint
func (m model) convert() (*entity.Topic, error) {
	return &entity.Topic{
		ID:          m.ID,
		CourseID:    m.CourseID,
		Title:       m.Title,
		OrderNumber: m.OrderNumber,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}, nil
}

// nolint
func (mm models) convert() ([]*entity.Topic, error) {
	out := make([]*entity.Topic, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
