package submissions

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/util"
	"github.com/samber/lo"
	"log"
)

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

func (s *service) ResetQuiz(ctx context.Context, userID, courseID, topicID string) error {
	passed, err := s.quizzesSubmissionRepository.CheckIsPassed(ctx, userID, topicID)
	if err != nil {
		return err
	}
	if !passed {
		return ErrNotPassed
	}
	err = s.quizzesSubmissionRepository.UpdatePassed(ctx, userID, courseID, topicID, false)
	return err
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
