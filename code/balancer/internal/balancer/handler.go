package balancer

import (
	"context"
	"encoding/json"
	"net/http"

	"ttb-service/internal/utils"
)

type Handler struct {
	Service BalancerService
}

// ApplyAlgorithmHandler - обработчик для применения алгоритма балансировки.
func (h *Handler) ApplyAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
	var req BalancingAlgorithmRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	result, err := h.Service.ApplyAlgorithm(context.Background(), req.AlgorithmType, req.Parameters)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to apply algorithm")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, result)
}

// ProcessEventHandler - обработчик для обработки события.
func (h *Handler) ProcessEventHandler(w http.ResponseWriter, r *http.Request) {
	var event PolicyEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid event payload")
		return
	}
	if err := h.Service.ProcessEvent(context.Background(), event); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to process event")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "processed"})
}

// SendEventHandler - обработчик для отправки события.
func (h *Handler) SendEventHandler(w http.ResponseWriter, r *http.Request) {
	var event PolicyEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid event payload")
		return
	}
	if err := h.Service.SendEvent(context.Background(), event); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send event")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "sent"})
}
