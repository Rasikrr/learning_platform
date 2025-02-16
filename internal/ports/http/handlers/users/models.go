package users

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"time"
)

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

type User struct {
	ID        string           `json:"id"`
	Name      *string          `json:"name"`
	LastName  *string          `json:"last_name"`
	Email     string           `json:"email"`
	Role      enum.AccountRole `json:"role"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt *time.Time       `json:"deleted_at"`
}

type userResponse struct {
	User User `json:"user"`
}

func convertUserResponse(user *entity.User) userResponse {
	return userResponse{
		User: convertUser(user),
	}
}

func convertUser(user *entity.User) User {
	return User{
		ID:        user.ID.String(),
		Name:      user.Name,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.AccountRole,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
