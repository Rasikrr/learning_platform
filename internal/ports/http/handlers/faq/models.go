package faq

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

//go:generate easyjson -all models.go

type postAnswerRequest struct {
	QuestionID string `json:"question_id"`
	Body       string `json:"body"`
}

type postQuestionRequest struct {
	Title      string  `json:"title"`
	Body       string  `json:"body"`
	CategoryID string  `json:"category_id"`
	ImageURL   *string `json:"image_url"`
}

type getAnswersRequest struct {
	QuestionID string
	Limit      int
	Offset     int
}

type getAnswersResponse struct {
	Answers []answer `json:"answers"`
}

type getQuestionsRequest struct {
	CategoryIDs []string `json:"category_ids"`
	Limit       int      `json:"limit"`
	Offset      int      `json:"offset"`
}

func (r getQuestionsRequest) ToParams() *entity.GetQuestionsParams {
	return &entity.GetQuestionsParams{
		CategoryIDs: r.CategoryIDs,
		Limit:       r.Limit,
		Offset:      r.Offset,
	}
}

type answer struct {
	ID         string    `json:"id"`
	QuestionID string    `json:"question_id"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Author     user      `json:"author"`
}

type user struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	LastName    *string `json:"last_name"`
	Email       string  `json:"email"`
	AccountRole string  `json:"account_role"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type getQuestionResponse struct {
	Question question `json:"question"`
}
type getQuestionsResponse struct {
	Questions []question `json:"questions"`
}

type question struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	ImageURL  *string  `json:"image_url"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Author    user     `json:"author"`
	Category  category `json:"category"`
}
type category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func convertToCategory(c *entity.QuestionCategory) category {
	return category{
		ID:        c.ID.String(),
		Name:      c.Name,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}
func convertToGetQuestionsResponse(questions []*entity.Question) getQuestionsResponse {
	out := make([]question, len(questions))
	for i, q := range questions {
		out[i] = convertToQuestion(q)
	}
	return getQuestionsResponse{
		Questions: out,
	}
}

func convertToQuestion(q *entity.Question) question {
	return question{
		ID:        q.ID.String(),
		Title:     q.Title,
		Body:      q.Body,
		ImageURL:  q.ImageURL,
		CreatedAt: q.CreatedAt.String(),
		UpdatedAt: q.UpdatedAt.String(),
		Author:    convertToUser(q.Author),
		Category:  convertToCategory(q.Category),
	}
}
func convertToQuestionResponse(q *entity.Question) getQuestionResponse {
	return getQuestionResponse{
		Question: convertToQuestion(q),
	}
}

func convertToUser(u *entity.User) user {
	return user{
		ID:          u.ID.String(),
		Name:        u.Name,
		LastName:    u.LastName,
		Email:       u.Email,
		AccountRole: u.AccountRole.String(),
		CreatedAt:   u.CreatedAt.String(),
		UpdatedAt:   u.UpdatedAt.String(),
	}
}

func convertToAnswer(a *entity.Answer) answer {
	return answer{
		ID:         a.ID.String(),
		QuestionID: a.Question.ID.String(),
		Body:       a.Body,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
		Author:     convertToUser(a.Author),
	}
}

func convertToGetAnswersResponse(answers []*entity.Answer) getAnswersResponse {
	out := make([]answer, len(answers))
	for i, a := range answers {
		out[i] = convertToAnswer(a)
	}
	return getAnswersResponse{
		Answers: out,
	}
}

func (r *getAnswersRequest) GetQueryParameters(req *http.Request) error {
	r.QuestionID = req.URL.Query().Get("question_id")
	limit := req.URL.Query().Get("limit")
	offset := req.URL.Query().Get("offset")

	var err error
	r.Limit, err = strconv.Atoi(limit)
	if err != nil {
		r.Limit = 15
	}
	r.Offset, err = strconv.Atoi(offset)
	if err != nil {
		r.Offset = 0
	}
	return nil
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
		ImageURL:  r.ImageURL,
		Title:     r.Title,
		Body:      r.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
