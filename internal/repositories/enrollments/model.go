// nolint
package enrollments

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	CourseID  uuid.UUID
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type models []model

// nolint
func (m model) convert() (*entity.Enrollment, error) {
	status, err := enum.CourseProgressString(m.Status)
	if err != nil {
		return nil, err
	}
	return &entity.Enrollment{
		ID:     m.ID,
		UserID: m.UserID,
		Course: entity.Course{
			ID: m.CourseID,
		},
		Status:    status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.CreatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.Enrollment, error) {
	out := make([]*entity.Enrollment, len(mm))
	for i, m := range mm {
		res, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}
