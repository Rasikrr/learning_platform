// nolint: revive, stylecheck, unused, unparam
package test_cases

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"time"
)

type model struct {
	ID             string
	TaskID         string
	Input          string
	ExpectedOutput string
	IsPublic       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type models []model

func (m *model) convert() (*entity.TestCase, error) {
	return &entity.TestCase{
		ID:             m.ID,
		TaskID:         m.TaskID,
		Input:          m.Input,
		ExpectedOutput: m.ExpectedOutput,
		IsPublic:       m.IsPublic,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}, nil
}

func (mm models) convert() ([]*entity.TestCase, error) {
	out := make([]*entity.TestCase, len(mm))
	for i, m := range mm {
		v, err := m.convert()
		if err != nil {
			return nil, err
		}
		out[i] = v
	}
	return out, nil
}
