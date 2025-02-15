package faq

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary post question
// @Description post question
// @Tags FAQ
// @Accept json
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param  request body postQuestionRequest true "request"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/faq/questions/post [post]
func (c *Controller) postQuestion(w http.ResponseWriter, r *http.Request) {
	session, err := api.GetSession(r.Context())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var req postQuestionRequest
	if err = api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	question, err := req.ToEntity(session)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()
	if err := c.faqService.PostQuestion(ctx, question); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
