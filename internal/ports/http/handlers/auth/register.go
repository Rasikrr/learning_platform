package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Register user
// @Description Register user with email and password and confirm password. Then send confirmation code to user email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "user credentials"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/auth/register [post]
func (c *Controller) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.authService.Register(r.Context(), req.Email, req.Password, req.PasswordConfirm); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
