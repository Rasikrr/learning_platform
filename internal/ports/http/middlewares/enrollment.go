package middlewares

import (
	"errors"
	"github.com/Rasikrr/learning_platform/internal/services/enrollments"
	"log"
	"net/http"
)

type EnrollMiddleware struct {
	enrollService enrollments.Service
}

func NewEnrollMiddleware(enrollService enrollments.Service) *EnrollMiddleware {
	return &EnrollMiddleware{
		enrollService: enrollService,
	}
}

func (m *EnrollMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("entering enroll middleware")
		session, err := GetSession(r.Context())
		if err != nil {
			SendError(w, http.StatusInternalServerError, err)
			return
		}
		courseID := r.URL.Query().Get("course_id")
		if courseID == "" {
			courseID = r.PathValue("course_id")
			if courseID == "" {
				SendError(w, http.StatusBadRequest, errors.New("course id is empty"))
				return
			}
		}
		enrolled, err := m.enrollService.CheckEnrollment(r.Context(), session.UserID.String(), courseID)
		if err != nil {
			SendError(w, http.StatusInternalServerError, err)
			return
		}
		if !enrolled {
			SendError(w, http.StatusBadRequest, errors.New("user is not enrolled in course"))
			return
		}
		next(w, r)
	}
}
