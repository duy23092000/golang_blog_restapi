package handler

import (
	"encoding/json"
	"net/http"
)

// responseJSON makes the response with payload as json format
func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// responseError makes the error response with payload as json format
func responseError(w http.ResponseWriter, code int, message string) {
	responseJSON(w, code, map[string]string{"error": message})
}
