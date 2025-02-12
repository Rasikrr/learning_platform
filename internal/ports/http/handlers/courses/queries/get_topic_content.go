package queries

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get course topic content
// @Description Get course topic content by topic id
// @Tags courses
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id path string true "course id"
// @Param topic_id path string true "topic id"
// @Success 200 {object} entity.TopicContent "Success"
// @Router /api/v1/courses/{course_id}/topic/{topic_id}/content [get]
func (c *Controller) getCourseTopicContent(w http.ResponseWriter, r *http.Request) {
	var req getTopicContentRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()

	content, err := c.coursesService.GetContentByTopicID(ctx, req.TopicID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, content, http.StatusOK)
}
