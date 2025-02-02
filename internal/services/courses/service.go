package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/categories"
	"github.com/Rasikrr/learning_platform/internal/repositories/content"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
	"github.com/Rasikrr/learning_platform/internal/repositories/quizzes"
	"github.com/Rasikrr/learning_platform/internal/repositories/tasks"
	"github.com/Rasikrr/learning_platform/internal/repositories/topics"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
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
	contentRepository    content.Repository
}

func NewService(
	coursesRepository courses.Repository,
	categoriesRepository categories.Repository,
	topicsRepository topics.Repository,
	quizzesRepository quizzes.Repository,
	tasksRepository tasks.Repository,
	contentRepository content.Repository,
) Service {
	return &service{
		coursesRepository:    coursesRepository,
		categoriesRepository: categoriesRepository,
		topicsRepository:     topicsRepository,
		quizzesRepository:    quizzesRepository,
		tasksRepository:      tasksRepository,
		contentRepository:    contentRepository,
	}
}

func (s *service) GetCoursesByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error) {
	courses, err := s.coursesRepository.GetByParams(ctx, params)
	if err != nil {
		return nil, err
	}
	catsIDs := lo.Map(courses, func(c *entity.Course, _ int) uuid.UUID {
		return c.CategoryID
	})
	categories, err := s.categoriesRepository.GetByIDs(ctx, catsIDs)
	if err != nil {
		return nil, err
	}
	categoriesMap := lo.SliceToMap(categories, func(t *entity.Category) (uuid.UUID, *entity.Category) {
		return t.ID, t
	})

	for _, c := range courses {
		category, ok := categoriesMap[c.CategoryID]
		if ok {
			c.Category = *category
		}
	}
	return courses, err
}

func (s *service) GetCourseByID(ctx context.Context, id string) (*entity.Course, error) {
	course, err := s.coursesRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = s.mergeCourse(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *service) GetAllCategories(ctx context.Context) ([]*entity.Category, error) {
	return s.categoriesRepository.GetAll(ctx)
}

func (s *service) mergeCourse(ctx context.Context, course *entity.Course) error {
	category, err := s.categoriesRepository.GetByID(ctx, course.CategoryID)
	if err != nil {
		return err
	}

	topics, err := s.topicsRepository.GetByCourseID(ctx, course.ID)
	if err != nil {
		return err
	}

	topicIDs := lo.Map(topics, func(t *entity.Topic, _ int) uuid.UUID {
		return t.ID
	})
	var (
		g       errgroup.Group
		content []*entity.TopicContent
		quizzes []*entity.Quiz
		tasks   []*entity.PracticalTask
	)
	g.Go(func() error {
		content, err = s.contentRepository.GetByTopicIDs(ctx, topicIDs)
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		quizzes, err = s.quizzesRepository.GetByTopicIDs(ctx, topicIDs)
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		tasks, err = s.tasksRepository.GetByTopicIDs(ctx, topicIDs)
		if err != nil {
			return err
		}
		return nil
	})

	if err = g.Wait(); err != nil {
		return err
	}

	topicMap := lo.SliceToMap(topics, func(t *entity.Topic) (uuid.UUID, *entity.Topic) {
		return t.ID, t
	})

	for _, c := range content {
		topic, ok := topicMap[c.TopicID]
		if ok {
			topic.Content = c
		}
	}
	for _, q := range quizzes {
		topic, ok := topicMap[q.TopicID]
		if ok {
			topic.Quizzes = append(topic.Quizzes, q)
		}
	}
	for _, p := range tasks {
		topic, ok := topicMap[p.TopicID]
		if ok {
			topic.PracticalTasks = append(topic.PracticalTasks, p)
		}
	}

	course.Category = *category
	course.Topics = topics

	return nil
}
