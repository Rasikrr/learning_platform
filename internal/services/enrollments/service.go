package enrollments

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
	"github.com/Rasikrr/learning_platform/internal/repositories/enrollments"
	"github.com/samber/lo"
)

type Service interface {
	GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error)
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

func (s *service) GetUserEnrollments(ctx context.Context, userID string) ([]*entity.Enrollment, error) {
	enrollments, err := s.enrollmentsRepository.GetUserEnrollments(ctx, userID)
	if err != nil {
		return nil, err
	}
	return s.mergeEnrollments(ctx, enrollments)
}

func (s *service) mergeEnrollments(ctx context.Context, enrollments []*entity.Enrollment) ([]*entity.Enrollment, error) {
	courses, err := s.coursesRepository.GetByIDs(ctx, lo.Map(enrollments, func(enrollment *entity.Enrollment, _ int) string {
		return enrollment.Course.ID.String()
	}))
	if err != nil {
		return nil, err
	}
	coursesMap := lo.SliceToMap(courses, func(course *entity.Course) (string, *entity.Course) {
		return course.ID.String(), course
	})
	for _, enrollment := range enrollments {
		enrollment.Course = *coursesMap[enrollment.Course.ID.String()]
	}
	return enrollments, nil
}
