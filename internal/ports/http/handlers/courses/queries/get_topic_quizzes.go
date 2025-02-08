package queries

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary get quizzes by topic id
// @Description get quizzes list by topic id
// @Tags courses
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param course_id query string true "course id"
// @Param topic_id query string true "topic id"
// @Success 200 {object} []entity.Quiz "Success"
// @Router /api/v1/courses/topic/quizzes [get]
func (c *Controller) getCourseTopicQuizzes(w http.ResponseWriter, r *http.Request) {
	var req getTopicQuizzesRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()

	quizzes, err := c.coursesService.GetQuizzesByTopicID(ctx, req.TopicID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, quizzes, http.StatusOK)
}
