package handlers

import (
	"encoding/json"
	"net/http"

	"ttb-service/internal/events"
	"ttb-service/internal/utils"
)

// EventHandler для обработки событий
type EventHandler struct {
	eventPublisher events.EventPublisher
	logger         *utils.Logger
}

// NewEventHandler создает новый обработчик для событий
func NewEventHandler(eventPublisher events.EventPublisher, logger *utils.Logger) *EventHandler {
	return &EventHandler{
		eventPublisher: eventPublisher,
		logger:         logger,
	}
}

// HandleEvent обрабатывает входящее событие
func (h *EventHandler) HandleEvent(w http.ResponseWriter, r *http.Request) {
	var event events.PolicyEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.logger.Error("Error decoding event body: " + err.Error())
		http.Error(w, "Invalid event format", http.StatusBadRequest)
		return
	}

	if err := h.eventPublisher.Publish(event.ToEvent()); err != nil {
		h.logger.Error("Error publishing event: " + err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
