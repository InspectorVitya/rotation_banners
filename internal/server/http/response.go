package httpserver

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(w http.ResponseWriter, statusCode int, msg string) error {
	msgWithErr := errorResponse{Message: msg}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(msgWithErr)
	if err != nil {
		return err
	}
	return nil
}
