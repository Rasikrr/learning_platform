package users

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/users"
)

type Service interface {
	DeleteUser(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
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

func (s *service) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return s.usersRepository.GetByID(ctx, id)
}

func (s *service) UpdateUser(ctx context.Context, params *entity.UpdateUserParams) error {
	return s.usersRepository.Update(ctx, params)
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	return s.usersRepository.Delete(ctx, id)
}
