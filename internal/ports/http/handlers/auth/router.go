package auth

import (
	authS "github.com/Rasikrr/learning_platform/internal/services/auth"

	"net/http"
)

type Controller struct {
	authService authS.Service
}

func NewController(authService authS.Service) *Controller {
	return &Controller{
		authService: authService,
	}
}

func (c *Controller) Init(r *http.ServeMux) {
	r.HandleFunc("POST /auth/login", c.Login)
	r.HandleFunc("POST/auth/register", c.Register)
	r.HandleFunc("POST /auth/logout", c.Logout)
	r.HandleFunc("POST /auth/confirm", c.ConfirmRegister)
}
