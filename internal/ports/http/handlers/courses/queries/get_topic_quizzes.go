package queries

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary get quizzes by topic id
// @Description get quizzes list by topic id. If quiz is passed, then answers will be returned
// @Tags quizzes
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id path string true "course id"
// @Param topic_id path string true "topic id"
// @Success 200 {object} getTopicQuizzesResponse "Success"
// @Router /api/v1/courses/{course_id}/topic/{topic_id}/quizzes [get]
func (c *Controller) getCourseTopicQuizzes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var req getTopicQuizzesRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}

	quizzes, passed, err := c.coursesService.GetQuizzesByTopicID(ctx, session.UserID.String(), req.TopicID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToGetTopicQuizzesResponse(quizzes, passed), http.StatusOK)
}
