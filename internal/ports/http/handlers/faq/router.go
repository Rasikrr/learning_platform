package faq

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	"github.com/Rasikrr/learning_platform/internal/services/faq"
	"net/http"
)

type Controller struct {
	faqService faq.Service
	m          *middlewares.AuthMiddleware
}

func NewController(authMiddleware *middlewares.AuthMiddleware, faqService faq.Service) *Controller {
	return &Controller{
		faqService: faqService,
		m:          authMiddleware,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("GET /api/v1/faq/answers", c.getAnswers)
	r.HandleFunc("GET /api/v1/faq/questions/{id}", c.getQuestion)
	r.HandleFunc("POST /api/v1/faq/questions", c.getQuestions)
	r.HandleFunc("GET /api/v1/faq/categories", c.getCategories)

	r.HandleFunc("POST /api/v1/faq/questions/post", c.m.Handle(c.postQuestion))
	r.HandleFunc("POST /api/v1/faq/answers/post", c.m.Handle(c.postAnswer))
}
