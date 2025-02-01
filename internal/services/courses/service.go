package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
)

type Service interface {
	GetCourses(ctx context.Context, limit, offset int) error
}

type service struct {
	coursesRepository courses.Repository
}

func NewService(coursesRepository courses.Repository) Service {
	return &service{
		coursesRepository: coursesRepository,
	}
}

func (s *service) GetCourses(ctx context.Context, limit, offset int) error {
	return s.coursesRepository.GetCourses(ctx, limit, offset)
}
