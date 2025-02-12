package courses

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID          uuid.UUID
	Title       string
	ImageURL    *string
	CategoryID  uuid.UUID
	Description string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type models []model

func (m model) convert() (*entity.Course, error) {
	category := entity.Category{
		ID: m.CategoryID,
	}
	return &entity.Course{
		ID:          m.ID,
		Title:       m.Title,
		ImageURL:    m.ImageURL,
		Category:    category,
		Description: m.Description,
		CreatedBy:   m.CreatedBy,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.Course, error) {
	out := make([]*entity.Course, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
