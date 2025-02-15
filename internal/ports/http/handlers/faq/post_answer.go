package faq

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary post answer
// @Description post answer
// @Tags FAQ
// @Accept json
// @Produce json
// @Security     BearerAuth
// @param Authorization header string true "Authorization token"
// @Param  request body postAnswerRequest true "request"
// @Success 200 {object} api.EmptySuccessResponse "Success"
// @Router /api/v1/faq/answers/post [post]
func (c *Controller) postAnswer(w http.ResponseWriter, r *http.Request) {
	session, err := api.GetSession(r.Context())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var req postAnswerRequest
	if err = api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	answer, err := req.ToEntity(session)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	ctx := r.Context()
	if err = c.faqService.PostAnswer(ctx, answer); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
