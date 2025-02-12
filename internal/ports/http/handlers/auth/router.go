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
	r.HandleFunc("POST /api/v1/auth/login", c.login)
	r.HandleFunc("POST /api/v1/auth/register", c.register)
	r.HandleFunc("POST /api/v1/auth/logout", c.logout)
	r.HandleFunc("POST /api/v1/auth/confirm", c.confirmRegister)
	r.HandleFunc("POST /api/v1/auth/refresh", c.refreshHandler)
	r.HandleFunc("POST /api/v1/auth/reset_password", c.resetPassword)
	r.HandleFunc("POST /api/v1/auth/confirm_reset_password", c.confirmResetPassword)
}
