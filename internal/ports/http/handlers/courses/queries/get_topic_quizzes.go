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
// @Param id path string true "topic id"
// @Success 200 {object} []entity.Quiz "Success"
// @Router /api/v1/courses/topic/{id}/quizzes [get]
func (c *Controller) getCourseTopicQuizzes(w http.ResponseWriter, r *http.Request) {
	topicID := r.PathValue("id")
	if topicID == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("topic id is empty"))
		return
	}
	ctx := r.Context()
	quizzes, err := c.coursesService.GetQuizzesByTopicID(ctx, topicID)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, quizzes, http.StatusOK)
}
