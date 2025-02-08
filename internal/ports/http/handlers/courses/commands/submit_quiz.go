package commands

import (
	"errors"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) SubmitQuiz(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	session, err := api.GetSession(ctx)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	var req submitQuizRequest
	if err = api.GetData(r, &req); err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	enrolled, err := c.enrollService.CheckEnrollment(ctx, session.UserID.String(), req.CourseID)
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	if !enrolled {
		api.SendError(w, http.StatusBadRequest, errors.New("user is not enrolled in course"))
		return
	}
	err = c.courseService.SubmitQuiz(ctx, req.CourseID, req.TopicID, req.Answers.ToEntities())
	if err != nil {
		api.SendError(w, http.StatusBadRequest, err)
		return
	}
	api.SendData(w, api.NewEmptySuccessResponse(), http.StatusOK)
}
