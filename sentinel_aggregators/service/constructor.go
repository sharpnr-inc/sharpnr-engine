package sentinelAggregatorService

import (
	"context"
	"fmt"

	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"

	"github.com/google/uuid"
)

// SentinelAggregatorServiceInterface aggregates every module's mod_resources
// into one catalog for the frontend policy builder: which resources exist, what
// actions/attributes they support, and their live instances. Exposed through
// the orchestrator; it makes no authorization decisions itself.
type SentinelAggregatorServiceInterface interface {
	// ListResources returns a summary of every registered resource type.
	ListResources(ctx context.Context, orgID, emplID uuid.UUID) ([]ResourceSummaryView, error)

	// GetResourceDetails returns one resource's full builder palette: its
	// actions, attributes and the org's live instances.
	GetResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) (ResourceDetailsView, error)
}

type sentinelAggregatorService struct {
	// ordered as registered at bootstrap, so the FE sees a stable catalog order
	descriptors []sentinelAggregatorContracts.ResourceDescriptor
	providers   map[string]sentinelAggregatorContracts.ResourceProvider // resource type -> owning provider
}

// Compile-time guard: the concrete service must satisfy the interface.
var _ SentinelAggregatorServiceInterface = (*sentinelAggregatorService)(nil)

// NewSentinelAggregatorService indexes the given providers by the resource
// types they declare. Two providers claiming the same Cedar entity type is a
// wiring bug, so it panics at bootstrap rather than serving an ambiguous
// catalog.
func NewSentinelAggregatorService(providers ...sentinelAggregatorContracts.ResourceProvider) *sentinelAggregatorService {
	s := &sentinelAggregatorService{
		providers: make(map[string]sentinelAggregatorContracts.ResourceProvider),
	}
	for _, provider := range providers {
		for _, descriptor := range provider.Resources() {
			if _, dup := s.providers[descriptor.Type]; dup {
				panic(fmt.Sprintf("sentinel aggregator: resource type %q registered twice", descriptor.Type))
			}
			s.providers[descriptor.Type] = provider
			s.descriptors = append(s.descriptors, descriptor)
		}
	}
	return s
}
