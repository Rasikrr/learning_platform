package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Name      string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(name, lastName, email, password string) *User {
	return &User{
		ID:        uuid.New(),
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
