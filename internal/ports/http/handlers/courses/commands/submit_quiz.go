package commands

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary submit quiz
// @Description submit quiz
// @Tags courses
// @Accept json
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id path string true "course id"
// @Param topic_id path string true "topic id"
// @Param  request body submitQuizRequest true "request"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/courses/{course_id}/topic/{topic_id}/quiz/submit [post]
func (c *Controller) submitQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var courseInfo courseAndTopic
	if err := api.GetData(r, &courseInfo); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	var req submitQuizRequest
	if err = api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	err = c.submissionService.SubmitQuiz(ctx, session.UserID.String(), courseInfo.CourseID, courseInfo.TopicID, req.Answers.ToEntities())
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
