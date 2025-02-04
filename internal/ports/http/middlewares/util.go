package middlewares

import (
	"encoding/json"
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
