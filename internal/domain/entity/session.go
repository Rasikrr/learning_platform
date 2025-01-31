package entity

import (
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
)

type Session struct {
	UserID uuid.UUID
	Email  string
	Role   enum.AccountRole
	Claims map[string]any
}

func (s *Session) AccountRole() enum.AccountRole {
	return s.Role
}
