package queries

import (
	"errors"
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
	session, err := api.GetSession(r.Context())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var req getTopicQuizzesRequest
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
	quizzes, err := c.coursesService.GetQuizzesByTopicID(ctx, req.TopicID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, quizzes, http.StatusOK)
}
