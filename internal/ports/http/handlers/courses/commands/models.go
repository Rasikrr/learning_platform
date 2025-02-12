package commands

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"net/http"
	"time"
)

//go:generate easyjson -all models.go

type courseAndTopic struct {
	CourseID string
	TopicID  string
	TaskID   string
}
type submitQuizRequest struct {
	Answers answers `json:"answers"`
}

type answerQuiz struct {
	QuestionID string `json:"question_id"`
	Answer     []bool `json:"answer"`
}

type submitTaskRequest struct {
	Input string `json:"input"`
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
	req.TaskID = r.PathValue("task_id")
	return nil
}

func (req submitTaskRequest) convert(c courseAndTopic, session *entity.Session) *entity.TaskSubmission {
	return &entity.TaskSubmission{
		TaskID:    c.TaskID,
		UserID:    session.UserID.String(),
		Input:     req.Input,
		CreatedAt: time.Now(),
	}
}
