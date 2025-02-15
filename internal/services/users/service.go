package users

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/users"
)

type Service interface {
	UpdateUser(ctx context.Context, params *entity.UpdateUserParams) error
}

type service struct {
	usersRepository users.Repository
}

func NewService(usersRepository users.Repository) Service {
	return &service{
		usersRepository: usersRepository,
	}
}

func (s *service) UpdateUser(ctx context.Context, params *entity.UpdateUserParams) error {
	return s.usersRepository.Update(ctx, params)
}
