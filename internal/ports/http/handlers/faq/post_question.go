package faq

import (
	"fmt"
	"github.com/Rasikrr/learning_platform/api"
	"net/http"
)

func (c *Controller) postQuestion(w http.ResponseWriter, r *http.Request) {
	session, err := api.GetSession(r.Context())
	if err != nil {
		api.SendError(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(session)
}
