package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) resetPassword(w http.ResponseWriter, r *http.Request) {
	var req ResetPasswordRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	if err := c.authService.ResetPassword(ctx, req.Email, req.Password, req.PasswordConfirm); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
