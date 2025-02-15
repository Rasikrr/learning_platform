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
	r.HandleFunc("POST /api/v1/admin/auth/register", c.register)
	r.HandleFunc("POST /api/v1/admin/auth/confirm", c.confirmRegister)
}
