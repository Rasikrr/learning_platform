package submissions

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
)

var (
	ErrCompilationError = errors.New("compilation error")
	ErrOutputNotEqual   = errors.New("output is not equal to expected output")
)

func (s *service) SubmitTask(ctx context.Context, submission *entity.TaskSubmission) (string, error) {
	task, err := s.tasksRepository.GetByID(ctx, submission.TaskID)
	if err != nil {
		return "", err
	}
	taskPassed, err := s.tasksSubmissionsRepository.CheckIsPassed(ctx, submission.UserID, task.ID.String())
	if err != nil {
		return "", err
	}
	if taskPassed {
		return "", ErrAlreadyPassed
	}
	if task.TestCases {
		return s.runTestsCases(ctx, task, submission)
	}
	var (
		out        string
		executeErr error
		passed     = true
	)
	output, err := s.taskExecutorClient.ExecuteCode(ctx, submission.Input, task.Language)
	if err != nil {
		out = output
		executeErr = err
		passed = false
	}
	if task.ExpectedOutput != nil && output != *task.ExpectedOutput {
		out = output
		executeErr = ErrOutputNotEqual
		passed = false
	}
	if executeErr != nil {
		executeErrString := executeErr.Error()
		submission.Error = &executeErrString
	}
	submission.Passed = passed
	if err = s.tasksSubmissionsRepository.Create(ctx, submission); err != nil {
		return "", err
	}
	return out, executeErr
}

func (s *service) runTestsCases(ctx context.Context, task *entity.PracticalTask, submission *entity.TaskSubmission) (string, error) {
	testCases, err := s.testCasesRepository.GetByTaskID(ctx, task.ID.String())
	if err != nil {
		return "", err
	}
	var (
		executeErr error
		passed     = true
		output     string
	)
	for _, testCase := range testCases {
		output, executeErr = s.taskExecutorClient.ExecuteTestCase(ctx, submission.Input, testCase.Input, task.Language)
		if executeErr != nil {
			executeErrString := executeErr.Error()
			submission.Error = &executeErrString
			passed = false
			break
		}
		if output != testCase.ExpectedOutput {
			executeErr = ErrOutputNotEqual
			executeErrString := executeErr.Error()
			submission.Error = &executeErrString
			executeErr = ErrOutputNotEqual
			passed = false
			break
		}
	}
	submission.Passed = passed
	if err = s.tasksSubmissionsRepository.Create(ctx, submission); err != nil {
		return "", err
	}
	return output, executeErr
}

func (s *service) ExecuteTask(ctx context.Context, input, taskID string) (string, error) {
	task, err := s.tasksRepository.GetByID(ctx, taskID)
	if err != nil {
		return "", err
	}
	return s.taskExecutorClient.ExecuteCode(ctx, input, task.Language)
}

func (s *service) ResetTask(ctx context.Context, userID, taskID string) error {
	passed, err := s.tasksSubmissionsRepository.CheckIsPassed(ctx, userID, taskID)
	if err != nil {
		return err
	}
	if !passed {
		return ErrNotPassed
	}
	err = s.tasksSubmissionsRepository.DeleteByUserAndTaskID(ctx, userID, taskID)
	return err
}
