package commands

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	coursesS "github.com/Rasikrr/learning_platform/internal/services/courses"
	"github.com/Rasikrr/learning_platform/internal/services/enrollments"
	submissionS "github.com/Rasikrr/learning_platform/internal/services/submissions"
	"net/http"
)

type Controller struct {
	courseService     coursesS.Service
	enrollService     enrollments.Service
	submissionService submissionS.Service

	m *middlewares.AuthMiddleware
}

func NewController(
	courseService coursesS.Service,
	m *middlewares.AuthMiddleware,
	submissionService submissionS.Service,
	enrollService enrollments.Service,
) *Controller {
	return &Controller{
		courseService:     courseService,
		enrollService:     enrollService,
		submissionService: submissionService,
		m:                 m,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.Handle("POST /api/v1/courses/submit_quiz", c.m.Handle(c.SubmitQuiz))
}
