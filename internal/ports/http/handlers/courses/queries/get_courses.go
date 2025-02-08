package queries

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get courses
// @Description Get courses for catalog
// @Tags courses
// @Produce json
// @Param request body getCoursesListRequest true "request"
// @Success 200 {object} getCoursesListResponse "Success"
// @Router /api/v1/courses [post]
func (c *Controller) getCourses(w http.ResponseWriter, r *http.Request) {
	var req getCoursesListRequest
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
	api.SendData(w, convertToGetCoursesListResponse(courses), http.StatusOK)
}
