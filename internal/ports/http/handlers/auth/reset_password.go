package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Reset password
// @Description Reset user password with email and password and confirm password. Then send confirmation code to user email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ResetPasswordRequest true "user credentials"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/auth/reset_password [post]
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
