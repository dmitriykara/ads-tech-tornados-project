package balancer

type PolicyEvent struct {
	EventType string                 `json:"eventType"`
	Details   map[string]interface{} `json:"details"`
}

type BalancingAlgorithmRequest struct {
	AlgorithmType string                 `json:"algorithmType"`
	Parameters    map[string]interface{} `json:"parameters"`
}

type BalancingResult struct {
	Success bool   `json:"success"`
	Details string `json:"details"`
}
