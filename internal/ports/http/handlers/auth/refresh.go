package auth

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

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
