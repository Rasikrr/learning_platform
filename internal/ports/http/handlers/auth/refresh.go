package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Refresh token
// @Description Refresh user tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RefreshRequest true "user's refresh token"
// @Success 200 {object} Auth "Success"
// @Router /api/v1/auth/refresh [post]
func (c *Controller) refreshHandler(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest

	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}

	tokens, err := c.authService.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToModel(tokens), http.StatusOK)
}
