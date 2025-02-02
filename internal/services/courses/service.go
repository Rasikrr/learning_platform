package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/categories"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
	"github.com/Rasikrr/learning_platform/internal/repositories/quizzes"
	"github.com/Rasikrr/learning_platform/internal/repositories/tasks"
	"github.com/Rasikrr/learning_platform/internal/repositories/topics"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetCoursesByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error)
	GetCourseByID(ctx context.Context, id string) (*entity.Course, error)
	GetAllCategories(ctx context.Context) ([]*entity.Category, error)
}

type service struct {
	coursesRepository    courses.Repository
	categoriesRepository categories.Repository
	topicsRepository     topics.Repository
	quizzesRepository    quizzes.Repository
	tasksRepository      tasks.Repository
}

func NewService(
	coursesRepository courses.Repository,
	categoriesRepository categories.Repository,
	topicsRepository topics.Repository,
	quizzesRepository quizzes.Repository,
	tasksRepository tasks.Repository,
) Service {
	return &service{
		coursesRepository:    coursesRepository,
		categoriesRepository: categoriesRepository,
		topicsRepository:     topicsRepository,
		quizzesRepository:    quizzesRepository,
		tasksRepository:      tasksRepository,
	}
}

func (s *service) GetCoursesByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error) {
	courses, err := s.coursesRepository.GetByParams(ctx, params)
	if err != nil {
		return nil, err
	}
	if err = s.mergeCourses(ctx, courses...); err != nil {
		return nil, err
	}
	return courses, nil
}

func (s *service) GetCourseByID(ctx context.Context, id string) (*entity.Course, error) {
	course, err := s.coursesRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = s.mergeCourses(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *service) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
	return s.categoriesRepository.GetAll(ctx)
}

func (s *service) mergeCourses(ctx context.Context, course ...*entity.Course) error {
	categoriesIDs := lo.Uniq(lo.Map(course, func(c *entity.Course, _ int) uuid.UUID {
		return c.CategoryID
	}))
	categories, err := s.categoriesRepository.GetByIDs(ctx, categoriesIDs)
	if err != nil {
		return err
	}

	topicsMap := lo.SliceToMap(categories, func(t *entity.Category) (uuid.UUID, *entity.Category) {
		return t.ID, t
	})
	for _, c := range course {
		if t, ok := topicsMap[c.CategoryID]; ok && t != nil {
			c.Category = *t
		}
	}

	return nil
}
