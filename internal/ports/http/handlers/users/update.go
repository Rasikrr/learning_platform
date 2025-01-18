package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req updateUserRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.usersService.UpdateName(r.Context(), req.Email, req.Name, req.Password); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
