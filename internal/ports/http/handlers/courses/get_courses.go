package courses

import (
	"github.com/Rasikrr/learning_platform/api"
	"log"
	"net/http"
)

func (c *Controller) getCourses(w http.ResponseWriter, r *http.Request) {
	var req getCoursesRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	log.Println("req", req)
	courses, err := c.coursesService.GetCourses(r.Context(), req.Limit, req.Offset)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, courses, http.StatusOK)
}
