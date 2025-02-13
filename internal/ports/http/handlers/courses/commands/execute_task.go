package commands

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// nolint
var (
	codeExample = "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"
)

// @Summary execute task
// @Description execute task
// @Tags courses
// @Accept json
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id path string true "course id"
// @Param topic_id path string true "topic id"
// @Param task_id path string true "task id"
// @Param  request body executeTaskRequest true "request"
// @Success 200 {object} executeTaskResponse "Success"
// @Router /api/v1/courses/{course_id}/topic/{topic_id}/task/{task_id}/execute [post]
func (c *Controller) executeTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var courseInfo courseAndTopic
	if err := api.GetData(r, &courseInfo); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}

	var req executeTaskRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}

	out, err := c.submissionService.ExecuteTask(ctx, req.Input, courseInfo.TaskID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToExecuteTaskResponse(out), http.StatusOK)
}
