package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithError отправляет JSON-ответ с ошибкой.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON отправляет JSON-ответ с указанными данными.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
