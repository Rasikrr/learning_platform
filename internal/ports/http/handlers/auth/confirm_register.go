package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) ConfirmRegister(w http.ResponseWriter, r *http.Request) {
	var req ConfirmRegisterRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.authService.ConfirmRegister(r.Context(), req.Email, req.Code); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
