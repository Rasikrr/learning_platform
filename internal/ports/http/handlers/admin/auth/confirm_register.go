package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Confirm admin register
// @Description Confirm admin user registration using confirmation code
// @Tags auth
// @Accept json
// @Produce json
// @Param request body confirmRegisterRequest true "user email and confirmation code"
// @Success 200 {object} Auth "Success"
// @Router /api/v1/admin/auth/confirm [post]
func (c *Controller) confirmRegister(w http.ResponseWriter, r *http.Request) {
	var req confirmRegisterRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	tokens, err := c.authService.ConfirmAdminRegister(r.Context(), req.Email, req.Code)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToModel(tokens), http.StatusOK)
}
