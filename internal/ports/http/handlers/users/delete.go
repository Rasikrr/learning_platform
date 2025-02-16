package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Delete my profile
// @Description Delete my profile
// @Tags users
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/users/me/delete [delete]
func (c *Controller) deleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	if err := c.usersService.DeleteUser(ctx, session.UserID.String()); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
