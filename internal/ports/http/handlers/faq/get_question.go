package faq

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary get question
// @Description post question by id
// @Tags FAQ
// @Accept json
// @Produce json
// @Param  id path string true "question id"
// @Success 200 {object} getQuestionResponse "Success"
// @Router /api/v1/faq/questions/{id} [get]
func (c *Controller) getQuestion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	questionID := r.PathValue("id")
	if questionID == "" {
		api.SendError(w, http.StatusBadRequest, errors.New("question id is required"))
		return
	}
	question, err := c.faqService.GetQuestion(ctx, questionID)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertToQuestionResponse(question), http.StatusOK)
}
