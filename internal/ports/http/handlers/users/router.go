package users

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	filesS "github.com/Rasikrr/learning_platform/internal/services/files"
	usersS "github.com/Rasikrr/learning_platform/internal/services/users"
	"net/http"
)

type Controller struct {
	usersService usersS.Service
	filesService filesS.Service
	m            *middlewares.AuthMiddleware
}

func NewController(
	usersService usersS.Service,
	filesService filesS.Service,
	m *middlewares.AuthMiddleware,
) *Controller {
	return &Controller{
		usersService: usersService,
		filesService: filesService,
		m:            m,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("GET /api/v1/users/me", c.m.Handle(c.getMyProfile))
	r.HandleFunc("GET /api/v1/users/{id}", c.m.Handle(c.getUser))
	r.HandleFunc("PUT /api/v1/users/update", c.m.Handle(c.updateUser))
	r.HandleFunc("DELETE /api/v1/users/me/delete", c.m.Handle(c.deleteUser))
}
