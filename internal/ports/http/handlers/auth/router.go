package auth

import (
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	authS "github.com/Rasikrr/learning_platform/internal/services/auth"

	"net/http"
)

type Controller struct {
	authService authS.Service
	m           *middlewares.AuthMiddleware
}

func NewController(authService authS.Service, m *middlewares.AuthMiddleware) *Controller {
	return &Controller{
		authService: authService,
		m:           m,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("POST /auth/login", c.login)
	r.HandleFunc("POST /auth/register", c.register)
	r.HandleFunc("POST /auth/logout", c.logout)
	r.HandleFunc("POST /auth/confirm", c.confirmRegister)
	r.HandleFunc("POST /auth/refresh", c.refreshHandler)
}
