package entity

import (
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID
	Name        string
	LastName    string
	Email       string
	Password    string
	AccountRole enum.AccountRole
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
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
