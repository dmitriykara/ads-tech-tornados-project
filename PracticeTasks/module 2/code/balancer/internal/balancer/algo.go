package balancer

import (
	"errors"
	"ttb-service/internal/utils"
)

// AlgorithmServiceImpl is a concrete implementation of the AlgorithmService interface.
type AlgorithmServiceImpl struct {
	logger *utils.Logger
}

// NewAlgorithmService creates a new instance of AlgorithmService.
func NewAlgorithmService(logger *utils.Logger) AlgorithmService {
	return &AlgorithmServiceImpl{
		logger: logger,
	}
}

// ApplyAlgorithm applies the balancing algorithm and returns the result.
func (s *AlgorithmServiceImpl) ApplyAlgorithm(request BalancingAlgorithmRequest) (BalancingResult, error) {
	// For demonstration purposes, we just handle a simple RoundRobin algorithm
	s.logger.Info("Applying balancing algorithm: " + request.AlgorithmType)
	switch request.AlgorithmType {
	case "RoundRobin":
		return BalancingResult{
			Success: true,
			Details: "RoundRobin algorithm applied successfully",
		}, nil
	case "LeastConnections":
		return BalancingResult{
			Success: true,
			Details: "LeastConnections algorithm applied successfully",
		}, nil
	case "Random":
		return BalancingResult{
			Success: true,
			Details: "Random algorithm applied successfully",
		}, nil
	default:
		return BalancingResult{}, errors.New("unknown algorithm type")
	}

}
