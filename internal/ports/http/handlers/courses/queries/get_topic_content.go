package queries

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

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
