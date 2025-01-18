package users

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entities"
	"github.com/Rasikrr/learning_platform/internal/repositories/users"
	"github.com/Rasikrr/learning_platform/internal/util"
	"strings"
)

type Service interface {
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	GetAll(ctx context.Context) ([]*entities.User, error)
	Create(ctx context.Context, name, email, password string) error
	Delete(ctx context.Context, email string) error
	UpdateName(ctx context.Context, email, name, password string) error
}

type service struct {
	userRepository users.Repository
}

func NewService(userRepository users.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetAll(ctx context.Context) ([]*entities.User, error) {
	out, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *service) Create(ctx context.Context, name, email, password string) error {
	u, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		if !strings.Contains(err.Error(), "email not found") {
			return err
		}
	}
	if u != nil {
		return errors.New("user with this email already exists")
	}
	user := entities.NewUser(name, email, password)
	if err := s.userRepository.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(ctx context.Context, email string) error {
	if err := s.userRepository.Delete(ctx, email); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateName(ctx context.Context, email, name, password string) error {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err := util.CheckPassword(user.Password, password); err != nil {
		return errors.New("incorrect password")
	}

	if err := s.userRepository.UpdateName(ctx, email, name); err != nil {
		return err
	}
	return nil
}
