package users

import "github.com/Rasikrr/learning_platform/internal/domain/entity"

//go:generate easyjson -all models.go

type updateUserRequest struct {
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
}

func (r updateUserRequest) ToEntity(session *entity.Session) *entity.UpdateUserParams {
	return &entity.UpdateUserParams{
		ID:       session.UserID.String(),
		Name:     r.Name,
		LastName: r.LastName,
	}
}
