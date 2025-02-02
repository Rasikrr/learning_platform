package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error)
	GetByID(ctx context.Context, id string) (*entity.Course, error)
	GetAllTopics(ctx context.Context) ([]*entity.Category, error)
}

type service struct {
	coursesRepository courses.Repository
}

func NewService(coursesRepository courses.Repository) Service {
	return &service{
		coursesRepository: coursesRepository,
	}
}

func (s *service) GetByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error) {
	courses, err := s.coursesRepository.GetByParams(ctx, params)
	if err != nil {
		return nil, err
	}
	if err = s.mergeCourses(ctx, courses...); err != nil {
		return nil, err
	}
	return courses, nil
}

func (s *service) GetByID(ctx context.Context, id string) (*entity.Course, error) {
	course, err := s.coursesRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = s.mergeCourses(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *service) GetAllTopics(ctx context.Context) ([]*entity.Category, error) {
	return s.coursesRepository.GetAllCategories(ctx)
}

func (s *service) mergeCourses(ctx context.Context, course ...*entity.Course) error {
	topicsIDs := lo.Uniq(lo.Map(course, func(c *entity.Course, _ int) uuid.UUID {
		return c.CategoryID
	}))
	topics, err := s.coursesRepository.GetCategoriesByIDs(ctx, topicsIDs)
	if err != nil {
		return err
	}
	topicsMap := lo.SliceToMap(topics, func(t *entity.Category) (uuid.UUID, *entity.Category) {
		return t.ID, t
	})
	for _, c := range course {
		if t, ok := topicsMap[c.CategoryID]; ok && t != nil {
			c.Category = *t
		}
	}
	return nil
}
