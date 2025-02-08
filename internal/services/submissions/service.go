package submissions

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/quizzes"
	quizzesSubmissionsR "github.com/Rasikrr/learning_platform/internal/repositories/quizzes_submissions"
	"github.com/Rasikrr/learning_platform/internal/util"
	"github.com/samber/lo"
	"log"
)

var (
	ErrNotFullyAnswered = errors.New("not fully answered")
	ErrNotCorrectAnswer = errors.New("not correct answer")
	ErrAlreadyPassed    = errors.New("already passed")
)

type Service interface {
	SubmitQuiz(ctx context.Context, userID, courseID, topicID string, answers []*entity.AnswerQuiz) error
}

type service struct {
	quizzesRepository           quizzes.Repository
	quizzesSubmissionRepository quizzesSubmissionsR.Repository
}

func NewService(
	quizzesRepository quizzes.Repository,
	quizzesSubmissionRepository quizzesSubmissionsR.Repository,
) Service {
	return &service{
		quizzesRepository:           quizzesRepository,
		quizzesSubmissionRepository: quizzesSubmissionRepository,
	}
}

func (s *service) SubmitQuiz(ctx context.Context, userID, courseID, topicID string, answers []*entity.AnswerQuiz) error {
	passed, err := s.quizzesSubmissionRepository.CheckIsPassed(ctx, userID, topicID)
	if err != nil {
		return err
	}
	log.Println("passed ", passed)
	if passed {
		return ErrAlreadyPassed
	}
	quizzes, err := s.quizzesRepository.GetByTopicID(ctx, topicID)
	if err != nil {
		return err
	}
	correct, checkErr := s.checkAnswers(quizzes, answers)
	err = s.quizzesSubmissionRepository.UpdatePassed(ctx, userID, courseID, topicID, correct)
	if err != nil {
		return err
	}
	return checkErr
}

func (s *service) checkAnswers(quizzes []*entity.Quiz, answers []*entity.AnswerQuiz) (bool, error) {
	answersMap := lo.SliceToMap(answers, func(a *entity.AnswerQuiz) (string, []bool) {
		return a.QuestionID, a.Answer
	})
	for _, quiz := range quizzes {
		ans, ok := answersMap[quiz.ID.String()]
		if !ok {
			return false, ErrNotFullyAnswered
		}
		if !util.SlicesEqual(ans, quiz.CorrectAnswers) {
			return false, ErrNotCorrectAnswer
		}
	}
	return true, nil
}
