package queries

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get course topic tasks
// @Description Get course topic tasks by topic id and. Also you must pass task order
// @Tags courses
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id query string true "course id"
// @Param topic_id query string true "topic id"
// @Param order query string true "number of task"
// @Success 200 {object} entity.PracticalTask "Success"
// @Router /api/v1/courses/topic/tasks [get]
func (c *Controller) getCourseTopicTasks(w http.ResponseWriter, r *http.Request) {
	var req getTopicTasksRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()

	tasks, err := c.coursesService.GetTasksByTopicIDAndOrderNum(ctx, req.TopicID, req.Order)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, tasks, http.StatusOK)
}
