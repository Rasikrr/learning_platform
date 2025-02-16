package courses

import (
	"context"
	coursesC "github.com/Rasikrr/learning_platform/internal/cache/courses"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/categories"
	"github.com/Rasikrr/learning_platform/internal/repositories/content"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
	"github.com/Rasikrr/learning_platform/internal/repositories/quizzes"
	quizzesSubmissionsR "github.com/Rasikrr/learning_platform/internal/repositories/quizzes_submissions"
	"github.com/Rasikrr/learning_platform/internal/repositories/tasks"
	"github.com/Rasikrr/learning_platform/internal/repositories/topics"
)

type Service interface {
	CreateCourse(ctx context.Context, params *entity.CreateCourseParams) error

	GetCoursesByParams(ctx context.Context, params *entity.GetCoursesParams) ([]*entity.Course, error)
	GetCourseByID(ctx context.Context, id string) (*entity.Course, error)

	GetAllCategories(ctx context.Context) ([]*entity.Category, error)

	GetContentByTopicID(ctx context.Context, topicID string) (*entity.TopicContent, error)

	GetQuizzesByTopicID(ctx context.Context, userID, topicID string) ([]*entity.Quiz, bool, error)

	GetTasksByTopicIDAndOrderNum(ctx context.Context, topicID string, order int) (*entity.PracticalTask, error)
}

type service struct {
	coursesRepository           courses.Repository
	categoriesRepository        categories.Repository
	topicsRepository            topics.Repository
	quizzesRepository           quizzes.Repository
	quizzesSubmissionRepository quizzesSubmissionsR.Repository
	tasksRepository             tasks.Repository
	contentRepository           content.Repository
	coursesCache                coursesC.Cache
}

func NewService(
	coursesRepository courses.Repository,
	categoriesRepository categories.Repository,
	topicsRepository topics.Repository,
	quizzesRepository quizzes.Repository,
	tasksRepository tasks.Repository,
	contentRepository content.Repository,
	quizzesSubmissionRepository quizzesSubmissionsR.Repository,
	coursesCache coursesC.Cache,
) Service {
	return &service{
		coursesRepository:           coursesRepository,
		categoriesRepository:        categoriesRepository,
		topicsRepository:            topicsRepository,
		quizzesRepository:           quizzesRepository,
		tasksRepository:             tasksRepository,
		contentRepository:           contentRepository,
		quizzesSubmissionRepository: quizzesSubmissionRepository,
		coursesCache:                coursesCache,
	}
}
