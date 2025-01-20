package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	tokens, err := c.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToModel(tokens), http.StatusOK)
}
