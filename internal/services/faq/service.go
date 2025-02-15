package faq

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/answers"
	questionCategories "github.com/Rasikrr/learning_platform/internal/repositories/question_categories"
	question "github.com/Rasikrr/learning_platform/internal/repositories/questions"
	"github.com/Rasikrr/learning_platform/internal/repositories/users"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	PostQuestion(ctx context.Context, question *entity.Question) error
	PostAnswer(ctx context.Context, answer *entity.Answer) error
	GetCategories(ctx context.Context) ([]*entity.QuestionCategory, error)
	GetAnswers(ctx context.Context, questionID string, limit, offset int) ([]*entity.Answer, error)
	GetQuestion(ctx context.Context, questionID string) (*entity.Question, error)
	GetQuestionsByParams(ctx context.Context, params *entity.GetQuestionsParams) ([]*entity.Question, error)
}

type service struct {
	questionsRepository  question.Repository
	categoriesRepository questionCategories.Repository
	answersRepository    answers.Repository
	usersRepository      users.Repository
}

func NewService(
	questionsRepository question.Repository,
	categoriesRepository questionCategories.Repository,
	answersRepository answers.Repository,
	usersRepository users.Repository,
) Service {
	return &service{
		questionsRepository:  questionsRepository,
		categoriesRepository: categoriesRepository,
		answersRepository:    answersRepository,
		usersRepository:      usersRepository,
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

func (s *service) GetCategories(ctx context.Context) ([]*entity.QuestionCategory, error) {
	return s.categoriesRepository.GetAll(ctx)
}

func (s *service) GetAnswers(ctx context.Context, questionID string, limit, offset int) ([]*entity.Answer, error) {
	answers, err := s.answersRepository.GetByQuestionID(ctx, questionID, limit, offset)
	if err != nil {
		return nil, err
	}
	return s.mergeAnswersUsers(ctx, answers...)
}

func (s *service) GetQuestion(ctx context.Context, questionID string) (*entity.Question, error) {
	uuid, err := uuid.Parse(questionID)
	if err != nil {
		return nil, err
	}
	question, err := s.questionsRepository.GetByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	merged, err := s.mergeQuestions(ctx, question)
	if err != nil {
		return nil, err
	}
	return merged[0], nil
}

func (s *service) GetQuestionsByParams(ctx context.Context, params *entity.GetQuestionsParams) ([]*entity.Question, error) {
	questions, err := s.questionsRepository.GetByParams(ctx, params)
	if err != nil {
		return nil, err
	}
	return s.mergeQuestions(ctx, questions...)
}

func (s *service) mergeQuestions(ctx context.Context, questions ...*entity.Question) ([]*entity.Question, error) {
	usersIDs := lo.Map(questions, func(q *entity.Question, _ int) string {
		return q.Author.ID.String()
	})
	var (
		users      []*entity.User
		categories []*entity.QuestionCategory
		g          errgroup.Group
	)
	catIDs := lo.Uniq(lo.Map(questions, func(q *entity.Question, _ int) string {
		return q.Category.ID.String()
	}))

	g.Go(func() error {
		var err error
		users, err = s.usersRepository.GetByIDs(ctx, usersIDs)
		return err
	})
	g.Go(func() error {
		var err error
		categories, err = s.categoriesRepository.GetByIDs(ctx, catIDs)
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	usersMap := lo.SliceToMap(users, func(u *entity.User) (string, *entity.User) {
		return u.ID.String(), u
	})
	categoriesMap := lo.SliceToMap(categories, func(c *entity.QuestionCategory) (string, *entity.QuestionCategory) {
		return c.ID.String(), c
	})
	for _, q := range questions {
		q.Author = usersMap[q.Author.ID.String()]
		q.Category = categoriesMap[q.Category.ID.String()]
	}
	return questions, nil
}

func (s *service) mergeAnswersUsers(ctx context.Context, answers ...*entity.Answer) ([]*entity.Answer, error) {
	usersIDs := lo.Map(answers, func(a *entity.Answer, _ int) string {
		return a.Author.ID.String()
	})
	users, err := s.usersRepository.GetByIDs(ctx, usersIDs)
	if err != nil {
		return nil, err
	}
	usersMap := lo.SliceToMap(users, func(u *entity.User) (string, *entity.User) {
		return u.ID.String(), u
	})
	for _, a := range answers {
		a.Author = usersMap[a.Author.ID.String()]
	}
	return answers, nil
}
