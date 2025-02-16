package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get my profile
// @Description Get my profile
// @Tags users
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Success 200 {object} userResponse "Success"
// @Router /api/v1/users/me [get]
func (c *Controller) getMyProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	user, err := c.usersService.GetByID(ctx, session.UserID.String())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertUserResponse(user), http.StatusOK)
}
