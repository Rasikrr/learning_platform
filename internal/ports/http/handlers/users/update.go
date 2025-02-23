package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Update user
// @Description Update user
// @Tags users
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param request body updateUserRequest true "request"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/users/update [put]
func (c *Controller) updateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var req updateUserRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	if err := c.usersService.UpdateUser(ctx, req.ToEntity(session)); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
