package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"net/http"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

// nolint: errcheck
func SendError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	r := ErrorResponse{
		Error:  err.Error(),
		Status: statusCode,
	}
	bb, _ := json.Marshal(r)
	w.Write(bb)
}

func GetSession(ctx context.Context) (*entity.Session, error) {
	token := ctx.Value(SessionKey)
	if token == nil {
		return nil, errors.New("session is empty")
	}
	s, ok := token.(*entity.Session)
	if !ok {
		return nil, errors.New("session is not Session")
	}
	return s, nil
}
