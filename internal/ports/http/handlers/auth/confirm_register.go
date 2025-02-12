package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Confirm register
// @Description Confirm user registration using confirmation code
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ConfirmRegisterRequest true "user email and confirmation code"
// @Success 200 {object} Auth "Success"
// @Router /api/v1/auth/confirm [post]
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
