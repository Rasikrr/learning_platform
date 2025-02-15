package faq

import (
	"github.com/Rasikrr/learning_platform/api"
	"log"
	"net/http"
)

// @Summary get answers
// @Description get answers for question
// @Tags FAQ
// @Produce json
// @Param question_id query string true "question id"
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 {object} getAnswersResponse "Success"
// @Router /api/v1/faq/answers [get]
func (c *Controller) getAnswers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req getAnswersRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	answers, err := c.faqService.GetAnswers(ctx, req.QuestionID, req.Limit, req.Offset)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	log.Println(answers)
	api.SendData(w, convertToGetAnswersResponse(answers), http.StatusOK)
}
