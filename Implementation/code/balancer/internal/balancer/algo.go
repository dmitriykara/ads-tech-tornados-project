package balancer

import (
	"sync"
	"ttb-service/internal/utils"
)

type ErrUnknownBalancingAlgorith struct {
	Message string
}

func (e ErrUnknownBalancingAlgorith) Error() string {
	return e.Message
}

// AlgorithmServiceImpl is a concrete implementation of the AlgorithmService interface.
type AlgorithmServiceImpl struct {
	logger       *utils.Logger
	currentValue BalancingAlgorithmRequest
	mu           sync.Mutex
}

// NewAlgorithmService creates a new instance of AlgorithmService.
func NewAlgorithmService(logger *utils.Logger) AlgorithmService {
	return &AlgorithmServiceImpl{
		logger: logger,
		currentValue: BalancingAlgorithmRequest{
			AlgorithmType: "LeastConnections",
			Parameters: map[string]interface{}{
				"version":        "1.0",
				"maxConnections": 10000,
				"healthCheck": map[string]interface{}{
					"interval":           30,
					"timeout":            5,
					"unhealthyThreshold": 3,
					"healthyThreshold":   2,
				},
			},
		},
		mu: sync.Mutex{},
	}
}

// ApplyAlgorithm applies the balancing algorithm and returns the result.
func (s *AlgorithmServiceImpl) ApplyAlgorithm(request BalancingAlgorithmRequest) (BalancingResult, error) {
	// For demonstration purposes, we just handle a simple RoundRobin algorithm
	s.logger.Info("Applying balancing algorithm: " + request.AlgorithmType)
	switch request.AlgorithmType {
	case "RoundRobin":
		s.apply(request)
		return BalancingResult{
			Success: true,
			Details: "RoundRobin algorithm applied successfully",
		}, nil
	case "LeastConnections":
		s.apply(request)
		return BalancingResult{
			Success: true,
			Details: "LeastConnections algorithm applied successfully",
		}, nil
	case "Random":
		s.apply(request)
		return BalancingResult{
			Success: true,
			Details: "Random algorithm applied successfully",
		}, nil
	default:
		return BalancingResult{}, ErrUnknownBalancingAlgorith{}
	}
}

// GetAlgorithm returns current balancing algorithm.
func (s *AlgorithmServiceImpl) GetAlgorithm() BalancingAlgorithmRequest {
	s.mu.Lock()
	algo := s.currentValue
	s.mu.Unlock()
	return algo
}

func (s *AlgorithmServiceImpl) apply(newValue BalancingAlgorithmRequest) {
	s.mu.Lock()
	s.currentValue = newValue
	s.mu.Unlock()
}
