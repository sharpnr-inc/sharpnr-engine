package sentinelAggregatorService

import (
	"context"

	"github.com/google/uuid"
)

type serviceConstructor struct {
	// add any fields needed for the service constructor
}

func NewSentinelAggregatorService() *serviceConstructor {
	return &serviceConstructor{}
}

// this methods will used in the other service to fetch the resources and their actions
// and resources attributes and their actions and types of actions
func (s *serviceConstructor) GetAllResources(ctx context.Context, orgID, emplID uuid.UUID) (string, error) {
	return "Sentinel Aggregator Service created successfully", nil
}

func (s *serviceConstructor) GetAllResourcesAttributes(ctx context.Context, orgID, emplID uuid.UUID, resourcesName string) (string, error) {
	return "Sentinel Aggregator Service created successfully", nil
}

func (s *serviceConstructor) GetAllResourcesActions(ctx context.Context, orgID, emplID uuid.UUID, resourcesName string) (string, error) {
	return "Sentinel Aggregator Service created successfully", nil
}
