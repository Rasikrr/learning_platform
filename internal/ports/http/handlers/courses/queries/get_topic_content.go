package queries

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get course topic content
// @Description Get course topic content by topic id
// @Tags courses
// @Produce json
// @Param id path string true "topic id"
// @Success 200 {object} entity.TopicContent "Success"
// @Router /api/v1/courses/topic/{id}/content [get]
func (c *Controller) getCourseTopicContent(w http.ResponseWriter, r *http.Request) {
	topicID := r.PathValue("id")
	if topicID == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("topic id is empty"))
		return
	}
	ctx := r.Context()
	content, err := c.coursesService.GetContentByTopicID(ctx, topicID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, content, http.StatusOK)
}
