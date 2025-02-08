package queries

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	"github.com/Rasikrr/learning_platform/internal/services/courses"
	"net/http"
)

type Controller struct {
	coursesService courses.Service
	e              *middlewares.EnrollMiddleware
	m              *middlewares.AuthMiddleware
}

func NewController(
	coursesService courses.Service,
	m *middlewares.AuthMiddleware,
	e *middlewares.EnrollMiddleware,
) *Controller {
	return &Controller{
		coursesService: coursesService,
		m:              m,
		e:              e,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	auth := c.m.Handle
	r.HandleFunc("POST /api/v1/courses", c.getCourses)
	r.HandleFunc("GET /api/v1/courses/categories", c.getCategories)
	r.HandleFunc("GET /api/v1/courses/{id}", c.getCourse)
	r.HandleFunc("GET /api/v1/courses/topic/content", auth(c.e.Handle(c.getCourseTopicContent)))
	r.HandleFunc("GET /api/v1/courses/topic/quizzes", auth(c.e.Handle(c.getCourseTopicQuizzes)))
	r.HandleFunc("GET /api/v1/courses/topic/tasks", auth(c.e.Handle(c.getCourseTopicTasks)))
}
