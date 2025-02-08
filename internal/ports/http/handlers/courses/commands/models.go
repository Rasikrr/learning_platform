package commands

import "github.com/Rasikrr/learning_platform/internal/domain/entity"

//go:generate easyjson -all models.go

type submitQuizRequest struct {
	CourseID string  `json:"course_id"`
	TopicID  string  `json:"topic_id"`
	Answers  answers `json:"answers"`
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
