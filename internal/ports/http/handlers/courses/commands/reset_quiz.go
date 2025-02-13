package commands

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary submit quiz
// @Description submit quiz
// @Tags quizzes
// @Accept json
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id path string true "course id"
// @Param topic_id path string true "topic id"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/courses/{course_id}/topic/{topic_id}/quiz/reset [put]
func (c *Controller) resetQuiz(w http.ResponseWriter, r *http.Request) {
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
	err = c.submissionService.ResetQuiz(ctx, session.UserID.String(), courseInfo.CourseID, courseInfo.TopicID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
