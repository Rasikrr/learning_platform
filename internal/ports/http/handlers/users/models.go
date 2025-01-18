package users

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entities"
	"time"
)

//go:generate easyjson -all models.go

type createUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type updateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type getAllUsersResponse struct {
	Result []User `json:"result"`
}

func convertEntity(e *entities.User) User {
	return User{
		Email:     e.Email,
		Name:      e.Name,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func newGetAllUsersResponse(users []*entities.User) getAllUsersResponse {
	out := make([]User, 0, len(users))
	for _, u := range users {
		out = append(out, convertEntity(u))
	}
	return getAllUsersResponse{
		Result: out,
	}
}
