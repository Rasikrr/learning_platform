package entity

import (
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"time"
)

//go:generate easyjson -all course.go

type Enrollment struct {
	ID        uuid.UUID           `json:"id"`
	UserID    uuid.UUID           `json:"user_id"`
	Course    Course              `json:"course"`
	Status    enum.CourseProgress `json:"status"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}
