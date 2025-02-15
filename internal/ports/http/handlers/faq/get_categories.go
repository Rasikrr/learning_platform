package faq

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary get categories
// @Description get question categories
// @Tags FAQ
// @Produce json
// @Success 200 {object} []entity.QuestionCategory "Success"
// @Router /api/v1/faq/categories [get]
func (c *Controller) getCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	categories, err := c.faqService.GetCategories(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, categories, http.StatusOK)
}
