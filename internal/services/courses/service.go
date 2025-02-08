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
)

type Service interface {
	GetCoursesByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error)
	GetCourseByID(ctx context.Context, id string) (*entity.Course, error)

	GetAllCategories(ctx context.Context) ([]*entity.Category, error)

	GetContentByTopicID(ctx context.Context, topicID string) (*entity.TopicContent, error)

	GetQuizzesByTopicID(ctx context.Context, topicID string) ([]*entity.Quiz, error)

	GetTasksByTopicIDAndOrderNum(ctx context.Context, topicID string, order int) (*entity.PracticalTask, error)

	SubmitQuiz(ctx context.Context, courseID string, topicID string, answers []*entity.AnswerQuiz) error
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
