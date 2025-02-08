package queries

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
	"strconv"
)

// @Summary Get course topic tasks
// @Description Get course topic tasks by topic id and. Also you must pass task order
// @Tags courses
// @Produce json
// @Param id path string true "topic id"
// @Param order path string true "order"
// @Success 200 {object} entity.PracticalTask "Success"
// @Router /api/v1/courses/topic/{id}/tasks/{order} [get]
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
