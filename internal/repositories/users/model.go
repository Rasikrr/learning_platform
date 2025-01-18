package users

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type models []*model

type model struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func convertModel(m *model) *entity.User {
	id, _ := uuid.Parse(m.ID)
	return &entity.User{
		ID:        id,
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func convertToModel(u *entity.User) *model {
	return &model{
		ID:        u.ID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func convertModels(ms models) []*entity.User {
	out := make([]*entity.User, len(ms))
	for i, m := range ms {
		out[i] = convertModel(m)
	}
	return out
}
