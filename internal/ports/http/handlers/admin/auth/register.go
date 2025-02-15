package auth

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"github.com/Rasikrr/learning_platform/internal/cache/auth"
	"net/http"
)

// @Summary Register admin user
// @Description Register admin user with email and password and confirm password. Then send confirmation code to user email
// @Tags auth
// @Accept json
// @Produce json
// @Param request body registerRequest true "user credentials"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/admin/auth/register [post]
func (c *Controller) register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.authService.Register(r.Context(), req.Email, req.Password, req.PasswordConfirm); err != nil {
		if !errors.Is(err, auth.ErrSpamDetected) {
			api.SendError(w, http.StatusBadRequest, err)
			return
		}
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
