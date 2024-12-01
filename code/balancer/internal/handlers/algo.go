package handlers

import (
	"encoding/json"
	"net/http"
	"ttb-service/internal/balancer"
	"ttb-service/internal/events"
	"ttb-service/internal/utils"
)

// AlgorithmBalancerHandler handles requests for applying balancing algorithms.
type AlgorithmBalancerHandler struct {
	algorithmService balancer.AlgorithmService
	logger           *utils.Logger
	eventPublisher   events.EventPublisher
}

// NewAlgorithmBalancerHandler creates a new instance of AlgorithmBalancerHandler.
func NewAlgorithmBalancerHandler(algorithmService balancer.AlgorithmService, logger *utils.Logger, eventPublisher events.EventPublisher) *AlgorithmBalancerHandler {
	return &AlgorithmBalancerHandler{
		algorithmService: algorithmService,
		logger:           logger,
		eventPublisher:   eventPublisher,
	}
}

// ApplyBalancerAlgorithm applies the requested balancing algorithm.
func (h *AlgorithmBalancerHandler) ApplyBalancerAlgorithm(w http.ResponseWriter, r *http.Request) {
	var req balancer.BalancingAlgorithmRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Error decoding request body: " + err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Apply the algorithm
	result, err := h.algorithmService.ApplyAlgorithm(req)
	if err != nil {
		h.logger.Error("Error applying balancing algorithm: " + err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Publish the event
	event := events.PolicyEvent{
		EventType: "PolicyApplied",
		Details:   result.Details,
	}
	if err := h.eventPublisher.Publish(event.ToEvent()); err != nil {
		h.logger.Error("Error publishing event: " + err.Error())
		http.Error(w, "Error publishing event", http.StatusInternalServerError)
		return
	}

	// Return the result
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		h.logger.Error("Error encoding response: " + err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
