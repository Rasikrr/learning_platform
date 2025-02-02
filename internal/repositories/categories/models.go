package categories

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedBy uuid.UUID `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type models []model

// nolint: unparam
func (m model) convert() (*entity.Category, error) {
	return &entity.Category{
		ID:        m.ID,
		Name:      m.Name,
		CreatedBy: m.CreatedBy,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.Category, error) {
	out := make([]*entity.Category, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
