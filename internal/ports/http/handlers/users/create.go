package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.usersService.Create(r.Context(), req.Name, req.Email, req.Password); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
