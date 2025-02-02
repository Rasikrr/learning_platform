package courses

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) getCourses(w http.ResponseWriter, r *http.Request) {
	var req getCoursesRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	params := req.toParams()
	courses, err := c.coursesService.GetCoursesByParams(ctx, params)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, courses, http.StatusOK)
}
