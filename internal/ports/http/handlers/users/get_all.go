package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.usersService.GetAll(r.Context())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, newGetAllUsersResponse(users), http.StatusOK)
}
