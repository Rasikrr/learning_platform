package queries

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	"github.com/Rasikrr/learning_platform/internal/services/courses"
	"github.com/Rasikrr/learning_platform/internal/services/enrollments"
	"net/http"
)

type Controller struct {
	coursesService     courses.Service
	enrollmentsService enrollments.Service
	m                  *middlewares.AuthMiddleware
}

func NewController(
	coursesService courses.Service,
	enrollmentsService enrollments.Service,
	m *middlewares.AuthMiddleware,
) *Controller {
	return &Controller{
		coursesService:     coursesService,
		enrollmentsService: enrollmentsService,
		m:                  m,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("POST /api/v1/courses", c.getCourses)
	r.HandleFunc("GET /api/v1/courses/categories", c.getCategories)
	r.HandleFunc("GET /api/v1/courses/{id}", c.getCourse)
	r.HandleFunc("GET /api/v1/courses/topic/content", c.m.Handle(c.getCourseTopicContent))
	r.HandleFunc("GET /api/v1/courses/topic/quizzes", c.m.Handle(c.getCourseTopicQuizzes))
	r.HandleFunc("GET /api/v1/courses/topic/tasks", c.m.Handle(c.getCourseTopicTasks))
}
