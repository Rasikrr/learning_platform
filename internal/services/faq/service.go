package faq

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/answers"
	questionCategories "github.com/Rasikrr/learning_platform/internal/repositories/question_categories"
	question "github.com/Rasikrr/learning_platform/internal/repositories/questions"
)

type Service interface {
	PostQuestion(ctx context.Context, question *entity.Question) error
	PostAnswer(ctx context.Context, answer *entity.Answer) error
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
	_, err := s.categoriesRepository.GetByID(ctx, question.Category.ID)
	if err != nil {
		return err
	}
	if err = s.questionsRepository.Create(ctx, question); err != nil {
		return err
	}
	return nil
}

func (s *service) PostAnswer(ctx context.Context, answer *entity.Answer) error {
	_, err := s.questionsRepository.GetByID(ctx, answer.Question.ID)
	if err != nil {
		return err
	}
	if err = s.answersRepository.Create(ctx, answer); err != nil {
		return err
	}
	return nil
}
