package submissions

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/clients/jdoodle"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/quizzes"
	quizzesSubmissionsR "github.com/Rasikrr/learning_platform/internal/repositories/quizzes_submissions"
	"github.com/Rasikrr/learning_platform/internal/repositories/tasks"
	tasksSubmissionsR "github.com/Rasikrr/learning_platform/internal/repositories/tasks_submissions"
	testCasesR "github.com/Rasikrr/learning_platform/internal/repositories/test_cases"
)

var (
	ErrNotFullyAnswered = errors.New("not fully answered")
	ErrNotCorrectAnswer = errors.New("not correct answer")
	ErrAlreadyPassed    = errors.New("already passed")
	ErrNotPassed        = errors.New("not passed")
)

type Service interface {
	SubmitQuiz(ctx context.Context, userID, courseID, topicID string, answers []*entity.AnswerQuiz) error
	ResetQuiz(ctx context.Context, userID, courseID, topicID string) error

	SubmitTask(ctx context.Context, submission *entity.TaskSubmission) (string, error)
	ExecuteTask(ctx context.Context, input, taskID string) (string, error)
	ResetTask(ctx context.Context, userID, taskID string) error
}

type service struct {
	quizzesRepository           quizzes.Repository
	quizzesSubmissionRepository quizzesSubmissionsR.Repository
	tasksSubmissionsRepository  tasksSubmissionsR.Repository
	tasksRepository             tasks.Repository
	testCasesRepository         testCasesR.Repository
	taskExecutorClient          jdoodle.Client
}

func NewService(
	quizzesRepository quizzes.Repository,
	quizzesSubmissionRepository quizzesSubmissionsR.Repository,
	tasksSubmissionsRepository tasksSubmissionsR.Repository,
	testCasesRepository testCasesR.Repository,
	tasksRepository tasks.Repository,
	taskExecutorClient jdoodle.Client,
) Service {
	return &service{
		quizzesRepository:           quizzesRepository,
		quizzesSubmissionRepository: quizzesSubmissionRepository,
		tasksSubmissionsRepository:  tasksSubmissionsRepository,
		testCasesRepository:         testCasesRepository,
		tasksRepository:             tasksRepository,
		taskExecutorClient:          taskExecutorClient,
	}
}
