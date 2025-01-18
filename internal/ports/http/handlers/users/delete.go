package users

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {

	email := r.PathValue("email")
	if email == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("email is required"))
		return
	}
	if err := c.usersService.Delete(r.Context(), email); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
