package queries

import (
	"errors"
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
	session, err := api.GetSession(r.Context())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}

	var req getTopicTasksRequest
	if err = api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()

	enrolled, err := c.enrollmentsService.CheckEnrollment(ctx, session.UserID.String(), req.CourseID)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	if !enrolled {
		api.SendError(w, http.StatusBadRequest, errors.New("user is not enrolled in course"))
		return
	}

	tasks, err := c.coursesService.GetTasksByTopicIDAndOrderNum(ctx, req.TopicID, req.Order)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, tasks, http.StatusOK)
}
