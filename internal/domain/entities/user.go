package entities

import (
	"github.com/Rasikrr/learning_platform/internal/util"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email, password string) *User {
	hPassword, _ := util.Hash(password)

	return &User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  hPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
