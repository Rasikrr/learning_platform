package auth

import (
	"github.com/Rasikrr/learning_platform/api"

	"net/http"
)

// @Summary Confirm password reset
// @Description Confirm user password reset using confirmation code
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ConfirmResetPasswordRequest true "user email and confirmation code"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/auth/confirm_reset_password [post]
func (c *Controller) confirmResetPassword(w http.ResponseWriter, r *http.Request) {
	var req ConfirmResetPasswordRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	if err := c.authService.ConfirmResetPassword(ctx, req.Email, req.Code); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
