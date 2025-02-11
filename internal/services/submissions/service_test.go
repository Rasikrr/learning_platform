// nolint: funlen
package submissions

import (
	"context"
	jdoodle "github.com/Rasikrr/learning_platform/internal/clients/mocks/jdoodle"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	quizzMock "github.com/Rasikrr/learning_platform/internal/repositories/mocks/quizzes"
	quizzesSubmissions "github.com/Rasikrr/learning_platform/internal/repositories/mocks/quizzes_submissions"
	tasks "github.com/Rasikrr/learning_platform/internal/repositories/mocks/tasks"
	tasksSubmissions "github.com/Rasikrr/learning_platform/internal/repositories/mocks/tasks_submissions"
	testCases "github.com/Rasikrr/learning_platform/internal/repositories/mocks/test_cases"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuizSubmission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizzesRepository := quizzMock.NewMockRepository(ctrl)
	mockQuizzesSubmissionRepository := quizzesSubmissions.NewMockRepository(ctrl)
	mockTaskSubmissionsRepository := tasksSubmissions.NewMockRepository(ctrl)
	mockTaskRepository := tasks.NewMockRepository(ctrl)
	mockTestCasesRepository := testCases.NewMockRepository(ctrl)

	ctx := context.Background()

	service := NewService(
		mockQuizzesRepository,
		mockQuizzesSubmissionRepository,
		mockTaskSubmissionsRepository,
		mockTestCasesRepository,
		mockTaskRepository,
		nil,
	)
	testCases := []quizTestCase{
		{
			Name:     "Test quiz submission success",
			UserID:   userID.String(),
			CourseID: courseID.String(),
			TopicID:  topicID.String(),
			Answers: []*entity.AnswerQuiz{
				{
					QuestionID: question1ID.String(),
					Answer:     []bool{true, false, false, false},
				},
				{
					QuestionID: question2ID.String(),
					Answer:     []bool{true, false, true, false},
				},
				{
					QuestionID: question3ID.String(),
					Answer:     []bool{false, false, false, true},
				},
			},
			ExpectedErr: nil,
		},
		{
			Name:     "Test quiz submission failed, not all answers correct",
			UserID:   userID.String(),
			CourseID: courseID.String(),
			TopicID:  topicID.String(),
			Answers: []*entity.AnswerQuiz{
				{
					QuestionID: question1ID.String(),
					Answer:     []bool{true, false, false, false},
				},
				{
					QuestionID: question2ID.String(),
					Answer:     []bool{true, false, true, false},
				},
				{
					QuestionID: question3ID.String(),
					Answer:     []bool{true, false, false, true},
				},
			},
			ExpectedErr: ErrNotCorrectAnswer,
		},
		{
			Name:     "Test quiz submission failed, not fully answered",
			UserID:   userID.String(),
			CourseID: courseID.String(),
			TopicID:  topicID.String(),
			Answers: []*entity.AnswerQuiz{
				{
					QuestionID: question1ID.String(),
					Answer:     []bool{true, false, false, false},
				},
				{
					QuestionID: question3ID.String(),
					Answer:     []bool{true, false, false, true},
				},
			},
			ExpectedErr: ErrNotFullyAnswered,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			mockQuizzesSubmissionRepository.EXPECT().CheckIsPassed(ctx, gomock.Any(), gomock.Any()).Return(false, nil)
			mockQuizzesRepository.EXPECT().GetByTopicID(ctx, gomock.Any()).Return(testQuizzes, nil)
			mockQuizzesSubmissionRepository.EXPECT().UpdatePassed(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			err := service.SubmitQuiz(ctx, testCase.UserID, testCase.CourseID, testCase.TopicID, testCase.Answers)
			assert.Equal(t, testCase.ExpectedErr, err)
		})
	}
}

func TestAlreadyPassed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizzesRepository := quizzMock.NewMockRepository(ctrl)
	mockQuizzesSubmissionRepository := quizzesSubmissions.NewMockRepository(ctrl)
	mockTaskSubmissionsRepository := tasksSubmissions.NewMockRepository(ctrl)
	mockTaskRepository := tasks.NewMockRepository(ctrl)
	mockTestCasesRepository := testCases.NewMockRepository(ctrl)

	ctx := context.Background()

	service := NewService(
		mockQuizzesRepository,
		mockQuizzesSubmissionRepository,
		mockTaskSubmissionsRepository,
		mockTestCasesRepository,
		mockTaskRepository,
		nil,
	)
	testCase := quizTestCase{
		Name:     "Test quiz submission, already passed",
		UserID:   userID.String(),
		CourseID: courseID.String(),
		TopicID:  topicID.String(),
		Answers: []*entity.AnswerQuiz{
			{
				QuestionID: question1ID.String(),
				Answer:     []bool{true, false, false, false},
			},
			{
				QuestionID: question2ID.String(),
				Answer:     []bool{true, false, true, false},
			},
			{
				QuestionID: question3ID.String(),
				Answer:     []bool{false, false, false, true},
			},
		},
		ExpectedErr: ErrAlreadyPassed,
	}

	t.Run(testCase.Name, func(t *testing.T) {
		mockQuizzesSubmissionRepository.EXPECT().CheckIsPassed(ctx, gomock.Any(), gomock.Any()).Return(true, nil)
		err := service.SubmitQuiz(ctx, testCase.UserID, testCase.CourseID, testCase.TopicID, testCase.Answers)
		assert.Equal(t, testCase.ExpectedErr, err)
	})
}

func TestTaskSubmission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizzesRepository := quizzMock.NewMockRepository(ctrl)
	mockQuizzesSubmissionRepository := quizzesSubmissions.NewMockRepository(ctrl)
	mockTaskSubmissionsRepository := tasksSubmissions.NewMockRepository(ctrl)
	mockTaskRepository := tasks.NewMockRepository(ctrl)
	mockTestCasesRepository := testCases.NewMockRepository(ctrl)
	mockJdoodleClient := jdoodle.NewMockClient(ctrl)

	ctx := context.Background()

	service := NewService(
		mockQuizzesRepository,
		mockQuizzesSubmissionRepository,
		mockTaskSubmissionsRepository,
		mockTestCasesRepository,
		mockTaskRepository,
		mockJdoodleClient,
	)

	testCases := []taskTestCase{
		{
			Name: "Test task submission success",
			Answers: &entity.TaskSubmission{
				TaskID: taskID.String(),
				UserID: userID.String(),
				Input:  taskSubmissionCode,
			},
			Output:      "3",
			ExpectedErr: nil,
		},
		{
			Name: "Test task submission failed, answer not equal to expected output",
			Answers: &entity.TaskSubmission{
				TaskID: taskID.String(),
				UserID: userID.String(),
				Input:  taskSubmissionCode,
			},
			Output:      "5",
			ExpectedErr: ErrOutputNotEqual,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			mockTaskRepository.EXPECT().GetByID(ctx, gomock.Any()).Return(testPracticalTask, nil)
			mockTaskSubmissionsRepository.EXPECT().CheckIsPassed(ctx, gomock.Any(), gomock.Any()).Return(false, nil)
			mockJdoodleClient.EXPECT().ExecuteCode(ctx, testCase.Answers.Input, enum.ProgrammingLanguageGo).Return(testCase.Output, nil)
			mockTaskSubmissionsRepository.EXPECT().Create(ctx, testCase.Answers).Return(nil)

			_, err := service.SubmitTask(ctx, testCase.Answers)
			assert.Equal(t, err, testCase.ExpectedErr)
		})
	}
}

func TestTaskSubmissionAlreadyPassed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuizzesRepository := quizzMock.NewMockRepository(ctrl)
	mockQuizzesSubmissionRepository := quizzesSubmissions.NewMockRepository(ctrl)
	mockTaskSubmissionsRepository := tasksSubmissions.NewMockRepository(ctrl)
	mockTaskRepository := tasks.NewMockRepository(ctrl)
	mockTestCasesRepository := testCases.NewMockRepository(ctrl)
	mockJdoodleClient := jdoodle.NewMockClient(ctrl)

	ctx := context.Background()

	service := NewService(
		mockQuizzesRepository,
		mockQuizzesSubmissionRepository,
		mockTaskSubmissionsRepository,
		mockTestCasesRepository,
		mockTaskRepository,
		mockJdoodleClient,
	)

	testCase := taskTestCase{

		Name: "Test task submission already passed",
		Answers: &entity.TaskSubmission{
			TaskID: taskID.String(),
			UserID: userID.String(),
			Input:  taskSubmissionCode,
		},
		Output:      "3",
		ExpectedErr: ErrAlreadyPassed,
	}

	t.Run(testCase.Name, func(t *testing.T) {
		t.Parallel()
		mockTaskRepository.EXPECT().GetByID(ctx, gomock.Any()).Return(testPracticalTask, nil)
		mockTaskSubmissionsRepository.EXPECT().CheckIsPassed(ctx, gomock.Any(), gomock.Any()).Return(true, nil)

		_, err := service.SubmitTask(ctx, testCase.Answers)
		assert.Equal(t, err, testCase.ExpectedErr)
	})
}

var (
	topicID, _     = uuid.Parse("10909c5e-987e-4302-b0dc-19a3307c54ec")
	courseID, _    = uuid.Parse("87da8c48-e419-49dd-9d4b-aca8e4fa76a2")
	question1ID, _ = uuid.Parse("2109a8f2-8cbc-4306-af4b-abe32bba763d")
	question2ID, _ = uuid.Parse("7441ee40-0315-4ab2-aa27-ef5ca8bf8250")
	question3ID, _ = uuid.Parse("18b92154-a4c8-4ebd-9e60-93da65051f93")
	userID, _      = uuid.Parse("e1e9e5f0-d7d8-4b0d-b1a1-e0e8d1d0c2c3")

	taskID, _ = uuid.Parse("3fbd5744-596a-4962-af51-c1390d25b034")

	testQuizzes = []*entity.Quiz{
		{
			ID:             question1ID,
			TopicID:        topicID,
			Question:       "question 1",
			Options:        []string{"option1", "option2", "option3", "option4"},
			CorrectAnswers: []bool{true, false, false, false},
			MultipleChoice: false,
		},
		{
			ID:             question2ID,
			TopicID:        topicID,
			Question:       "question 2",
			Options:        []string{"option1", "option2", "option3", "option4"},
			CorrectAnswers: []bool{true, false, true, false},
			MultipleChoice: true,
		},
		{
			ID:             question3ID,
			TopicID:        topicID,
			Question:       "question 3",
			Options:        []string{"option1", "option2", "option3", "option4"},
			CorrectAnswers: []bool{false, false, false, true},
			MultipleChoice: false,
		},
	}

	taskSubmissionCode = `package main
import (
    "fmt"
    "time"
)
func main() {
    fmt.Println(1+2)
}
`
	expectedOutput    = "3"
	testPracticalTask = &entity.PracticalTask{
		ID:             taskID,
		Description:    "simple go task, 1+2",
		ExpectedOutput: &expectedOutput,
		TestCases:      false,
		Language:       enum.ProgrammingLanguageGo,
	}
)

type quizTestCase struct {
	Name        string
	UserID      string
	CourseID    string
	TopicID     string
	Answers     []*entity.AnswerQuiz
	ExpectedErr error
}

type taskTestCase struct {
	Name        string
	Answers     *entity.TaskSubmission
	Output      string
	ExpectedErr error
}
