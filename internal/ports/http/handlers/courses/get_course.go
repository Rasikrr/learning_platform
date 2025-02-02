package courses

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) getCourse(w http.ResponseWriter, r *http.Request) {
	courseID := r.PathValue("id")
	if courseID == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("course id is empty"))
		return
	}
	ctx := r.Context()
	courses, err := c.coursesService.GetByID(ctx, courseID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, courses, http.StatusOK)
}
