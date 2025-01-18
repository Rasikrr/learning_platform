package users

import (
	usersS "github.com/Rasikrr/learning_platform/internal/services/users"
	"net/http"
)

type Controller interface {
	Init(r *http.ServeMux)
}

type controller struct {
	usersService usersS.Service
}

func New(usersService usersS.Service) Controller {
	return &controller{
		usersService: usersService,
	}
}

func (c *controller) Init(r *http.ServeMux) {
	r.HandleFunc("GET /users/{email}", c.GetUser)
	r.HandleFunc("GET /users/all", c.GetAllUsers)
	r.HandleFunc("POST /users/", c.CreateUser)
	r.HandleFunc("PUT /users/", c.UpdateUser)
	r.HandleFunc("DELETE /users/{email}", c.DeleteUser)
}
