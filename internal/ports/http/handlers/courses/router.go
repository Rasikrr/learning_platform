package courses

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
	r.HandleFunc("GET /courses", c.getCourses)
}
