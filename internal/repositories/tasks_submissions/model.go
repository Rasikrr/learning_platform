// nolint: revive, stylecheck, unused
package tasks_submissions

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"time"
)

type model struct {
	ID        string
	TaskID    string
	UserID    string
	Input     string
	Passed    bool
	Error     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type models []model

// nolint: unparam
func (m *model) convert() (*entity.TaskSubmission, error) {
	return &entity.TaskSubmission{
		ID:        m.ID,
		TaskID:    m.TaskID,
		UserID:    m.UserID,
		Input:     m.Input,
		Passed:    m.Passed,
		Error:     m.Error,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.TaskSubmission, error) {
	out := make([]*entity.TaskSubmission, 0, len(mm))
	for _, m := range mm {
		v, err := m.convert()
		if err != nil {
			return nil, err
		}
		if v != nil {
			out = append(out, v)
		}
	}
	return out, nil
}

func convert(sub *entity.TaskSubmission) *model {
	return &model{
		ID:        sub.ID,
		TaskID:    sub.TaskID,
		UserID:    sub.UserID,
		Input:     sub.Input,
		Passed:    sub.Passed,
		Error:     sub.Error,
		CreatedAt: sub.CreatedAt,
	}
}
