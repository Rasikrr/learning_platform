package faq

import (
	questionCategories "github.com/Rasikrr/learning_platform/internal/repositories/question_categories"
	question "github.com/Rasikrr/learning_platform/internal/repositories/questions"
)

type Service interface{}

type service struct {
	questionsRepository  question.Repository
	categoriesRepository questionCategories.Repository
	answersRepository    question.Repository
}

func NewService(
	questionsRepository question.Repository,
	categoriesRepository questionCategories.Repository,
	answersRepository question.Repository,
) Service {
	return &service{
		questionsRepository:  questionsRepository,
		categoriesRepository: categoriesRepository,
		answersRepository:    answersRepository,
	}
}
