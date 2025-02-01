package entity

import (
	"github.com/google/uuid"
	"time"
)

type Course struct {
	ID          uuid.UUID
	Title       string
	ImageURL    *string
	TopicID     uuid.UUID
	Description string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
