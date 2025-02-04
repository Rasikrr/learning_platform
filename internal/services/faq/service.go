package faq

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/answers"
	questionCategories "github.com/Rasikrr/learning_platform/internal/repositories/question_categories"
	question "github.com/Rasikrr/learning_platform/internal/repositories/questions"
)

type Service interface {
	PostQuestion(ctx context.Context, question *entity.Question) error
}

type service struct {
	questionsRepository  question.Repository
	categoriesRepository questionCategories.Repository
	answersRepository    answers.Repository
}

func NewService(
	questionsRepository question.Repository,
	categoriesRepository questionCategories.Repository,
	answersRepository answers.Repository,
) Service {
	return &service{
		questionsRepository:  questionsRepository,
		categoriesRepository: categoriesRepository,
		answersRepository:    answersRepository,
	}
}

func (s *service) PostQuestion(ctx context.Context, question *entity.Question) error {
	category, err := s.categoriesRepository.GetByID(ctx, question.ID)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("category not found")
	}
	if err := s.questionsRepository.Create(ctx, question); err != nil {
		return err
	}
	return nil
}
