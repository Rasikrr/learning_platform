package faq

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

//go:generate easyjson -all models.go

type postAnswerRequest struct {
	QuestionID string `json:"question_id"`
	Body       string `json:"body"`
}

type postQuestionRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryID string `json:"category_id"`
}

func (r postAnswerRequest) Validate() error {
	if len(r.Body) == 0 {
		return errors.New("body is required")
	}
	_, err := uuid.Parse(r.QuestionID)
	if err != nil {
		return errors.New("invalid question_id")
	}
	return nil
}

func (r postAnswerRequest) ToEntity(session *entity.Session) (*entity.Answer, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	questionID, err := uuid.Parse(r.QuestionID)
	if err != nil {
		return nil, errors.New("invalid question_id")
	}
	question := &entity.Question{
		ID: questionID,
	}
	answer := &entity.Answer{
		ID:       uuid.New(),
		Question: question,
		Body:     r.Body,
		Author:   api.GetUserFromSession(session),
	}
	return answer, nil
}

func (r postQuestionRequest) Validate() error {
	if len(r.Title) == 0 {
		return errors.New("title is required")
	}
	if len(r.Body) == 0 {
		return errors.New("body is required")
	}
	_, err := uuid.Parse(r.CategoryID)
	if err != nil {
		return errors.New("invalid category_id")
	}
	return nil
}

func (r postQuestionRequest) ToEntity(session *entity.Session) (*entity.Question, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}
	categoryID, err := uuid.Parse(r.CategoryID)
	if err != nil {
		return nil, err
	}
	category := &entity.QuestionCategory{
		ID: categoryID,
	}
	return &entity.Question{
		ID:        uuid.New(),
		Category:  category,
		Author:    api.GetUserFromSession(session),
		Title:     r.Title,
		Body:      r.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
