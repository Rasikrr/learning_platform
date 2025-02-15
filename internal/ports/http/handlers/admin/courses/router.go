package courses

import (
	coursesS "github.com/Rasikrr/learning_platform/internal/services/courses"
	"net/http"
)

type Controller struct {
	courseService coursesS.Service
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("POST /api/v1/admin/courses/create", c.create)
}
