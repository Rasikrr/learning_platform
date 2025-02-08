// nolint: revive, stylecheck, unused
package quizzes_submissions

import (
	"github.com/google/uuid"
	"time"
)

type model struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TopicID   uuid.UUID
	CourseID  uuid.UUID
	Passed    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type models []model
