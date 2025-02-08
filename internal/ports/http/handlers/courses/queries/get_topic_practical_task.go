package queries

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
	"strconv"
)

func (c *Controller) getCourseTopicTasks(w http.ResponseWriter, r *http.Request) {
	topicID, order := r.PathValue("id"), r.PathValue("order")
	if topicID == "" || order == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("topic id is empty"))
		return
	}
	orderNum, err := strconv.Atoi(order)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	tasks, err := c.coursesService.GetTasksByTopicIDAndOrderNum(ctx, topicID, orderNum)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, tasks, http.StatusOK)
}
