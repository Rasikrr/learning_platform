package courses

import (
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID          uuid.UUID
	Title       string
	ImageURL    *string
	TopicID     uuid.UUID
	Description string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type models []model
