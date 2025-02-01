package courses

import (
	"net/http"
	"strconv"
)

type getCoursesRequest struct {
	Limit  int
	Offset int
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
	r.Limit = limit
	r.Offset = offset
	return nil
}
