package queries

import (
	"github.com/Rasikrr/learning_platform/internal/services/courses"
	"net/http"
)

type Controller struct {
	coursesService courses.Service
}

func NewController(coursesService courses.Service) *Controller {
	return &Controller{
		coursesService: coursesService,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("POST /api/v1/courses", c.getCourses)
	r.HandleFunc("GET /api/v1/courses/categories", c.getCategories)
	r.HandleFunc("GET /api/v1/courses/{id}", c.getCourse)
	r.HandleFunc("GET /api/v1/courses/topic/{id}/content", c.getCourseTopicContent)
	r.HandleFunc("GET /api/v1/courses/topic/{id}/quizzes", c.getCourseTopicQuizzes)
	r.HandleFunc("GET /api/v1/courses/topic/{id}/tasks/{order}", c.getCourseTopicTasks)
}
