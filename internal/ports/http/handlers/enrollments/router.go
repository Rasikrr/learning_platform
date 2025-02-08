package commands

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	"github.com/Rasikrr/learning_platform/internal/services/enrollments"
	"net/http"
)

type Controller struct {
	enrollmentsService enrollments.Service
	m                  *middlewares.AuthMiddleware
}

func NewController(enrollmentsService enrollments.Service, m *middlewares.AuthMiddleware) *Controller {
	return &Controller{
		enrollmentsService: enrollmentsService,
		m:                  m,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("POST /api/v1/courses/enroll", c.m.Handle(c.enrollToCourse))
}
