package courses

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) getCategories(w http.ResponseWriter, r *http.Request) {
	topics, err := c.coursesService.GetAllTopics(r.Context())
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, topics, http.StatusOK)
}
