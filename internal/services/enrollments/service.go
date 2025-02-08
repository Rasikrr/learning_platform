package enrollments

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
	"github.com/Rasikrr/learning_platform/internal/repositories/enrollments"
)

type Service interface {
	Enroll(ctx context.Context, userID string, courseID string) error
	CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error)
}

type service struct {
	coursesRepository     courses.Repository
	enrollmentsRepository enrollments.Repository
}

func NewService(
	coursesRepository courses.Repository,
	enrollmentsRepository enrollments.Repository,
) Service {
	return &service{
		coursesRepository:     coursesRepository,
		enrollmentsRepository: enrollmentsRepository,
	}
}

func (s *service) Enroll(ctx context.Context, userID string, courseID string) error {
	course, err := s.coursesRepository.GetByID(ctx, courseID)
	if err != nil {
		return err
	}
	enrolled, err := s.CheckEnrollment(ctx, userID, course.ID.String())
	if err != nil {
		return err
	}
	if enrolled {
		return errors.New("user already enrolled")
	}
	return s.enrollmentsRepository.Enroll(ctx, userID, course.ID.String())
}

func (s *service) CheckEnrollment(ctx context.Context, userID string, courseID string) (bool, error) {
	exists, err := s.enrollmentsRepository.CheckByUserIDAndCourseID(ctx, userID, courseID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
