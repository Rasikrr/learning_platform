package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Login user
// @Description Login user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "user credentials"
// @Success 200 {object} Auth "Success"
// @Router /api/v1/auth/login [post]
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
