package users

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	if email == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}
	user, err := c.usersService.GetByEmail(r.Context(), email)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertEntity(user), http.StatusOK)
}
