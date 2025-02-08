package entity

import (
	"github.com/google/uuid"
	"time"
)

//go:generate easyjson -all faq.go

type Question struct {
	ID        uuid.UUID         `json:"id"`
	Category  *QuestionCategory `json:"category"`
	Title     string            `json:"title"`
	Body      string            `json:"body"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Author    *User             `json:"author"`
}

type QuestionCategory struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Answer struct {
	ID        uuid.UUID `json:"id"`
	Question  *Question `json:"question"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    *User     `json:"author"`
}
