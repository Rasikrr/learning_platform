package commands

import (
	"github.com/Rasikrr/learning_platform/api"
	"log"
	"net/http"
)

// @Summary submit task
// @Description submit task
// @Tags courses
// @Accept json
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id path string true "course id"
// @Param topic_id path string true "topic id"
// @Param task_id path string true "task id"
// @Param  request body submitTaskRequest true "request"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/courses/{course_id}/topic/{topic_id}/task/{task_id}/submit [post]
func (c *Controller) submitTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}

	var courseInfo courseAndTopic
	if err = api.GetData(r, &courseInfo); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}

	var req submitTaskRequest
	if err = api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	submission := req.convert(courseInfo, session)
	out, err := c.submissionService.SubmitTask(ctx, submission)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	log.Println("out", out)
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
