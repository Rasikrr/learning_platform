package commands

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	coursesS "github.com/Rasikrr/learning_platform/internal/services/courses"
	submissionS "github.com/Rasikrr/learning_platform/internal/services/submissions"
	"net/http"
)

type Controller struct {
	courseService     coursesS.Service
	submissionService submissionS.Service
	e                 *middlewares.EnrollMiddleware
	m                 *middlewares.AuthMiddleware
}

func NewController(
	courseService coursesS.Service,
	m *middlewares.AuthMiddleware,
	e *middlewares.EnrollMiddleware,
	submissionService submissionS.Service,
) *Controller {
	return &Controller{
		courseService:     courseService,
		submissionService: submissionService,
		m:                 m,
		e:                 e,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	auth := c.m.Handle
	r.Handle("POST /api/v1/courses/{course_id}/topic/{topic_id}/quiz/submit", auth(c.e.Handle(c.submitQuiz)))
	r.Handle("POST /api/v1/courses/{course_id}/topic/{topic_id}/quiz/reset", auth(c.e.Handle(c.resetQuiz)))
}
