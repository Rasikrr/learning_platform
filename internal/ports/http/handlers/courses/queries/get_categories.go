package queries

import (
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

// @Summary Get courses categories
// @Description Get courses categories
// @Tags courses
// @Produce json
// @Success 200 {object} getCategoriesListResponse "Success"
// @Router /api/v1/courses/categories [get]
func (c *Controller) getCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	categories, err := c.coursesService.GetAllCategories(ctx)
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, convertToGetCategoriesListResponse(categories), http.StatusOK)
}
