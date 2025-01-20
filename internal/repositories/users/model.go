package users

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"log"
	"time"
)

// nolint: unused
type models []model

type model struct {
	ID          string
	Name        string
	LastName    string
	Email       string
	Password    string
	AccountRole string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func convertModel(m model) (*entity.User, error) {
	log.Printf("%+v\n", m)
	role, err := enum.AccountRoleString(m.AccountRole)
	if err != nil {
		return nil, err
	}
	id, _ := uuid.Parse(m.ID)
	return &entity.User{
		ID:          id,
		Name:        m.Name,
		LastName:    m.LastName,
		Email:       m.Email,
		AccountRole: role,
		Password:    m.Password,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}, nil
}

func convertToModel(u *entity.User) model {
	return model{
		ID:          u.ID.String(),
		Name:        u.Name,
		LastName:    u.LastName,
		Email:       u.Email,
		Password:    u.Password,
		AccountRole: u.AccountRole.String(),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		DeletedAt:   u.DeletedAt,
	}
}

// nolint: unused
func convertModels(ms models) ([]*entity.User, error) {
	var err error
	out := make([]*entity.User, len(ms))
	for i, m := range ms {
		out[i], err = convertModel(m)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
