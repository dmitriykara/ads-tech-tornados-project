package handlers

import (
	"net/http"
)

// HealthCheckHandler для обработки запросов на проверку состояния
type HealthCheckHandler struct{}

// NewHealthCheckHandler создает новый обработчик для health check
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// CheckHealth проверяет здоровье сервиса
func (h *HealthCheckHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is healthy"))
}
