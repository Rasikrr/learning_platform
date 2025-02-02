package courses

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
)

//go:generate easyjson -all models.go

type getCoursesRequest struct {
	Limit         int      `json:"limit"`
	Offset        int      `json:"offset"`
	CategoriesIDs []string `json:"categories_ids"`
}

func (r *getCoursesRequest) toParams() *entity.GetCoursesParams {
	if r.Limit == 0 {
		r.Limit = 15
	}
	return &entity.GetCoursesParams{
		Limit:         r.Limit,
		Offset:        r.Offset,
		CategoriesIDs: r.CategoriesIDs,
	}
}
