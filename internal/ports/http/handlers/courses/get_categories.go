package courses

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) getCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	topics, err := c.coursesService.GetAllCategories(ctx)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, topics, http.StatusOK)
}
