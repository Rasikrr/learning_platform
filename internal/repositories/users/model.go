package users

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

// nolint: unused
type models []model

type model struct {
	ID        string
	Name      string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func convertModel(m model) *entity.User {
	id, _ := uuid.Parse(m.ID)
	return &entity.User{
		ID:        id,
		Name:      m.Name,
		LastName:  m.LastName,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: m.DeletedAt,
	}
}

func convertToModel(u *entity.User) model {
	return model{
		ID:        u.ID.String(),
		Name:      u.Name,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

// nolint: unused
func convertModels(ms models) []*entity.User {
	out := make([]*entity.User, len(ms))
	for i, m := range ms {
		out[i] = convertModel(m)
	}
	return out
}
