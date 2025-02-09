package commands

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"net/http"
)

//go:generate easyjson -all models.go

type courseAndTopic struct {
	CourseID string
	TopicID  string
}
type submitQuizRequest struct {
	Answers answers `json:"answers"`
}

type answerQuiz struct {
	QuestionID string `json:"question_id"`
	Answer     []bool `json:"answer"`
}

type answers []answerQuiz

func (a answers) ToEntities() []*entity.AnswerQuiz {
	out := make([]*entity.AnswerQuiz, len(a))
	for i, answer := range a {
		out[i] = answer.ToEntity()
	}
	return out
}

func (r answerQuiz) ToEntity() *entity.AnswerQuiz {
	return &entity.AnswerQuiz{
		QuestionID: r.QuestionID,
		Answer:     r.Answer,
	}
}

func (req *courseAndTopic) GetParameters(r *http.Request) error {
	req.CourseID = r.PathValue("course_id")
	req.TopicID = r.PathValue("topic_id")
	return nil
}
