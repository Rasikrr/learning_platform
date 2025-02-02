package entity

import (
	"github.com/google/uuid"
	"time"
)

//go:generate easyjson -all course.go

type Course struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ImageURL    *string   `json:"image_url"`
	CategoryID  uuid.UUID `json:"category_id"`
	Topic       Topic     `json:"topic"`
	Description string    `json:"description"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Topic struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
