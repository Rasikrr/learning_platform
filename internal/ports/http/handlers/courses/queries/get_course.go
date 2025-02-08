package queries

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get course by id
// @Description Get detailed course info with topics by id
// @Tags courses
// @Produce json
// @Param id path string true "course id"
// @Success 200 {object} getCourseDetailedResponse "Success"
// @Router /api/v1/courses/{id} [get]
func (c *Controller) getCourse(w http.ResponseWriter, r *http.Request) {
	courseID := r.PathValue("id")
	if courseID == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("course id is empty"))
		return
	}
	ctx := r.Context()
	courses, err := c.coursesService.GetCourseByID(ctx, courseID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToGetCourseDetailedResponse(courses), http.StatusOK)
}
