package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/ports/http/middlewares"
	"io"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

type ParametersGetter interface {
	GetParameters(r *http.Request) error
}

// nolint: errcheck
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	marshaller, ok := data.(json.Marshaler)
	if ok {
		bb, err := marshaller.MarshalJSON()
		if err != nil {
			SendError(w, http.StatusInternalServerError, err)
			return
		}
		w.Write(bb)
		return
	}
	bb, err := json.Marshal(data)
	if err != nil {
		SendError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(bb)
}

func GetData(r *http.Request, data interface{}) error {
	params, ok := data.(ParametersGetter)
	if ok {
		if err := params.GetParameters(r); err != nil {
			return err
		}
	}
	log.Println(params)
	defer r.Body.Close()
	bb, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(bb) > 0 {
		log.Println(bb)
		unmarshaller, ok := data.(json.Unmarshaler)
		if ok {
			return unmarshaller.UnmarshalJSON(bb)
		}
	}
	return nil
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
	token := ctx.Value(middlewares.SessionKey)
	if token == nil {
		return nil, errors.New("session is empty")
	}
	s, ok := token.(*entity.Session)
	if !ok {
		return nil, errors.New("session is not Session")
	}
	return s, nil
}

func GetUserFromSession(session *entity.Session) *entity.User {
	return &entity.User{
		ID:          session.UserID,
		Email:       session.Email,
		AccountRole: session.Role,
	}
}
