package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

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
	defer r.Body.Close()
	bb, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(bb) > 0 {
		unmarshaller, ok := data.(json.Unmarshaler)
		if ok {
			return unmarshaller.UnmarshalJSON(bb)
		}
		// TODO: get data from query params
	}
	return json.Unmarshal(bb, data)
}

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
