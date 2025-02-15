package faq

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary get questions
// @Description post questions by params (category_ids, limit, offset)
// @Tags FAQ
// @Accept json
// @Produce json
// @Param  request body getQuestionsRequest true "request"
// @Success 200 {object} getQuestionsResponse "Success"
// @Router /api/v1/faq/questions [post]
func (c *Controller) getQuestions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getQuestionsRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
	}
	params := req.ToParams()
	questions, err := c.faqService.GetQuestionsByParams(ctx, params)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, convertToGetQuestionsResponse(questions), http.StatusOK)
}
