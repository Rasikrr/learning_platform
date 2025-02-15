package courses

import (
	"github.com/Rasikrr/learning_platform/api"
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"net/http"
)

func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	if !session.Role.OneOf(enum.AccountRoleAdmin) {
		api.SendError(w, http.StatusForbidden, err)
		return
	}
	var req createCourseRequest
	if err := api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	params := req.ToEntity(session)
	if err = c.courseService.CreateCourse(ctx, params); err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
