package middlewares

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("internal server error")
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				SendError(w, http.StatusInternalServerError, ErrInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
