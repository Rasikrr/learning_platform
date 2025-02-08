package courses

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/util"
	"github.com/samber/lo"
)

var (
	ErrNotFullyAnswered = errors.New("not fully answered")
	ErrNotCorrectAnswer = errors.New("not correct answer")
)

func (s *service) SubmitQuiz(ctx context.Context, _ string, topicID string, answers []*entity.AnswerQuiz) error {
	quizzes, err := s.quizzesRepository.GetByTopicID(ctx, topicID)
	if err != nil {
		return err
	}
	answersMap := lo.SliceToMap(answers, func(a *entity.AnswerQuiz) (string, []bool) {
		return a.QuestionID, a.Answer
	})
	for _, quiz := range quizzes {
		ans, ok := answersMap[quiz.ID.String()]
		if !ok {
			return ErrNotFullyAnswered
		}
		if !util.SlicesEqual(ans, quiz.CorrectAnswers) {
			return ErrNotCorrectAnswer
		}
	}
	return nil
}
