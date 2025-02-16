package users

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get user
// @Description Get user by id
// @Tags users
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} userResponse "Success"
// @Router /api/v1/users/{id} [get]
func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.PathValue("id")

	user, err := c.usersService.GetByID(ctx, userID)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertUserResponse(user), http.StatusOK)
}
