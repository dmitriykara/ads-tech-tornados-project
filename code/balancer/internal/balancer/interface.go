package balancer

import "context"

// BalancerService - интерфейс логики балансировщика.
type BalancerService interface {
	ApplyAlgorithm(ctx context.Context, algorithmType string, params map[string]interface{}) (BalancingResult, error)
	ProcessEvent(ctx context.Context, event PolicyEvent) error
	SendEvent(ctx context.Context, event PolicyEvent) error
}

// BalancerRepository - интерфейс работы с базой данных.
type BalancerRepository interface {
	SaveBalancingResult(ctx context.Context, result BalancingResult) error
	GetAlgorithmParameters(ctx context.Context, algorithmType string) (map[string]interface{}, error)
}

// AlgorithmService defines the methods to apply the balancing algorithm.
type AlgorithmService interface {
	// ApplyAlgorithm applies the balancing algorithm based on the provided request.
	ApplyAlgorithm(request BalancingAlgorithmRequest) (BalancingResult, error)
}
