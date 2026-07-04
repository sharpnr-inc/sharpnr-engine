package sentinelAggregatorService

import (
	"context"
	"fmt"

	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"

	"github.com/google/uuid"
)

// ResourceSummaryView is the list shape of one protectable resource — enough
// for the FE to render the resource picker without the full palette.
type ResourceSummaryView struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// ResourceDetailsView is one resource's full policy-builder palette: the static
// descriptor (actions + attributes) plus the org's live instances.
type ResourceDetailsView struct {
	sentinelAggregatorContracts.ResourceDescriptor
	Instances []sentinelAggregatorContracts.ResourceInstance `json:"instances"`
}

// ListResources implements [SentinelAggregatorServiceInterface].
func (s *sentinelAggregatorService) ListResources(ctx context.Context, orgID, emplID uuid.UUID) ([]ResourceSummaryView, error) {
	out := make([]ResourceSummaryView, 0, len(s.descriptors))
	for _, d := range s.descriptors {
		out = append(out, ResourceSummaryView{
			Type:        d.Type,
			Label:       d.Label,
			Description: d.Description,
		})
	}
	return out, nil
}

// GetResourceDetails implements [SentinelAggregatorServiceInterface].
func (s *sentinelAggregatorService) GetResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) (ResourceDetailsView, error) {
	provider, ok := s.providers[resourceType]
	if !ok {
		return ResourceDetailsView{}, fmt.Errorf("unknown resource type %q", resourceType)
	}

	instances, err := provider.ListInstances(ctx, orgID, emplID, resourceType)
	if err != nil {
		return ResourceDetailsView{}, err
	}
	if instances == nil {
		instances = []sentinelAggregatorContracts.ResourceInstance{}
	}

	for _, d := range s.descriptors {
		if d.Type == resourceType {
			return ResourceDetailsView{ResourceDescriptor: d, Instances: instances}, nil
		}
	}
	// unreachable: providers and descriptors are indexed together
	return ResourceDetailsView{}, fmt.Errorf("unknown resource type %q", resourceType)
}
