package courses

import (
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"net/http"
	"strconv"
)

type getCoursesRequest struct {
	Limit    int
	Offset   int
	Category *string
}

func (r *getCoursesRequest) GetParameters(req *http.Request) error {
	limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
	if err != nil {
		limit = 15
	}
	offset, err := strconv.Atoi(req.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}
	category := req.URL.Query().Get("category")
	if category != "" {
		r.Category = &category
	}
	r.Limit = limit
	r.Offset = offset
	return nil
}

func (r *getCoursesRequest) toParams() *entity.GetCoursesParams {
	return &entity.GetCoursesParams{
		Limit:    r.Limit,
		Offset:   r.Offset,
		Category: r.Category,
	}
}
