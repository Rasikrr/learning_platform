package middlewares

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	authS "github.com/Rasikrr/learning_platform/internal/services/auth"
	"log"
	"net/http"
)

type sessionKeyType string

const (
	SessionKey sessionKeyType = "session"
	authHeader string         = "Authorization"
)

type AuthMiddleware struct {
	authService authS.Service
}

func NewAuthMiddleware(authService authS.Service) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("entering auth middleware")
		session, err := m.parseAuth(r)
		if err != nil {
			SendError(w, http.StatusUnauthorized, err)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), SessionKey, session))
		next(w, r)
	}
}

func (m *AuthMiddleware) parseAuth(r *http.Request) (*entity.Session, error) {
	token := r.Header.Get(authHeader)
	if token == "" {
		return nil, errors.New("authorization header is empty")
	}
	ctx := r.Context()
	session, err := m.authService.CheckToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return session, nil
}
