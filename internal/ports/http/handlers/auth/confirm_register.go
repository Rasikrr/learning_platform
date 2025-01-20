package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) confirmRegister(w http.ResponseWriter, r *http.Request) {
	var req ConfirmRegisterRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	tokens, err := c.authService.ConfirmRegister(r.Context(), req.Email, req.Code)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToModel(tokens), http.StatusOK)
}
