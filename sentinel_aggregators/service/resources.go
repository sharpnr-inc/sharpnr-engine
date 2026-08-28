package sentinelAggregatorService

import (
	"context"
	"errors"
	"fmt"

	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

	"github.com/google/uuid"
)

// ErrUnknownResourceType is returned when the requested resource type is not in
// the catalog. Controllers match it with errors.Is to answer 404 rather than 500
// — an unrecognised type in the URL is a client mistake, not a server fault.
var ErrUnknownResourceType = errors.New("sentinel aggregator: unknown resource type")

// ResourceSummaryView is the list shape of one protectable resource — enough
// for the FE to render the resource picker without the full palette.
type ResourceSummaryView struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

// ResourceDetailsView is one resource's full policy-builder palette: the static
// descriptor (actions + attributes) plus the org's live instances.
//
// HasInstances distinguishes "this type has no selectable instances, author
// type-level policies only" from "instances exist but this org has none" —
// both of which arrive as an empty Instances array.
type ResourceDetailsView struct {
	sentinelAggregatorContracts.ResourceDescriptor
	HasInstances bool                                           `json:"hasInstances"`
	Instances    []sentinelAggregatorContracts.ResourceInstance `json:"instances"`
}

/* ---------------------------------- <***> --------------------------------- */

// GetContextPalette implements [SentinelAggregatorServiceInterface]. The
// palette is static app vocabulary (defined next to the keys the engine
// stamps, in sentinel_contexts); orgID/emplID are carried for signature
// consistency and future per-org keys.
func (s *sentinelAggregatorService) GetContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error) {
	return sentinelContexts.BuilderPalette(), nil
}

/* ---------------------------------- <***> --------------------------------- */

// ListResources implements [SentinelAggregatorServiceInterface], in registration
// order so the catalog is stable across calls.
func (s *sentinelAggregatorService) ListResources(ctx context.Context, orgID, emplID uuid.UUID) ([]ResourceSummaryView, error) {
	out := make([]ResourceSummaryView, 0, len(s.order))
	for _, resourceType := range s.order {
		d := s.resources[resourceType].descriptor
		out = append(out, ResourceSummaryView{
			Type:        d.Type,
			Label:       d.Label,
			Description: d.Description,
		})
	}
	return out, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetResourceDetails implements [SentinelAggregatorServiceInterface].
//
// One map lookup resolves both the descriptor and the lister bound to this
// type, so no module is asked to dispatch on a resource type the aggregator has
// already resolved.
func (s *sentinelAggregatorService) GetResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) (ResourceDetailsView, error) {
	entry, ok := s.resources[resourceType]
	if !ok {
		return ResourceDetailsView{}, fmt.Errorf("%w: %q", ErrUnknownResourceType, resourceType)
	}

	view := ResourceDetailsView{
		ResourceDescriptor: entry.descriptor,
		HasInstances:       entry.list != nil,
		Instances:          []sentinelAggregatorContracts.ResourceInstance{},
	}
	if entry.list == nil {
		// No instances by design: policies about this type are type-level only.
		return view, nil
	}

	instances, err := entry.list(ctx, orgID, emplID)
	if err != nil {
		return ResourceDetailsView{}, fmt.Errorf("listing instances of %q: %w", resourceType, err)
	}
	if instances != nil {
		view.Instances = instances
	}
	return view, nil
}
