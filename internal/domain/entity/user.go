package entity

import (
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"time"
)

//go:generate easyjson -all user.go

type User struct {
	ID          uuid.UUID        `json:"id"`
	Name        *string          `json:"name"`
	LastName    *string          `json:"last_name"`
	Email       string           `json:"email"`
	Password    string           `json:"password"`
	AccountRole enum.AccountRole `json:"account_role"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   *time.Time       `json:"deleted_at"`
}

func NewUser(email, password string) *User {
	return &User{
		ID:          uuid.New(),
		Email:       email,
		Password:    password,
		AccountRole: enum.AccountRoleUser,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func NewAdminUser(email, password string) *User {
	return &User{
		ID:          uuid.New(),
		Email:       email,
		Password:    password,
		AccountRole: enum.AccountRoleAdmin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type UpdateUserParams struct {
	ID       string  `json:"id"`
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
}
