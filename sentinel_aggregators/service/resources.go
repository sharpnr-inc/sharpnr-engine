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

// ErrUnknownAction is returned by ValidateActions for an id no module declared.
var ErrUnknownAction = errors.New("sentinel aggregator: unknown action")

// ResourceSummaryView is the list shape — enough to render a resource picker.
type ResourceSummaryView struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category"`
	HousedIn    string `json:"housedIn,omitempty"`
}

// ResourceDetailsView is one resource's full palette plus a page of instances.
//
// HasInstances distinguishes "this type has no selectable instances, scope it
// through its container instead" from "instances exist but this org has none" —
// both of which arrive as an empty array.
type ResourceDetailsView struct {
	sentinelAggregatorContracts.ResourceDescriptor
	HasInstances bool                                           `json:"hasInstances"`
	Instances    []sentinelAggregatorContracts.ResourceInstance `json:"instances"`
	NextCursor   string                                         `json:"nextCursor,omitempty"`
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
			Category:    d.Category,
			HousedIn:    d.HousedIn,
		})
	}
	return out, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetCatalog implements [SentinelAggregatorServiceInterface]: every descriptor
// with its actions and attributes, in one response.
//
// The permission matrix needs actions for every resource at once. Composing that
// from ListResources plus one GetResourceDetails per type is a request per
// resource on a screen that opens constantly, and the details call also does
// instance listing the matrix does not want.
func (s *sentinelAggregatorService) GetCatalog(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorContracts.ResourceDescriptor, error) {
	out := make([]sentinelAggregatorContracts.ResourceDescriptor, 0, len(s.order))
	for _, resourceType := range s.order {
		d, err := s.describe(ctx, orgID, s.resources[resourceType])
		if err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}

/* ---------------------------------- <***> --------------------------------- */

// describe returns the descriptor with per-org value lists filled in.
//
// Resolvers run on the catalog read, never on the authorization path — the
// frontend fetches this once and caches it, so a query per attribute here is
// paid once a session rather than once a decision. A resolver that fails takes
// the whole catalog with it rather than serving an attribute with a silently
// empty picker, which would look like "this org has no cities" instead of
// "something is broken".
func (s *sentinelAggregatorService) describe(
	ctx context.Context, orgID uuid.UUID, entry registeredResource,
) (sentinelAggregatorContracts.ResourceDescriptor, error) {
	d := entry.descriptor
	if len(entry.resolveValues) == 0 {
		return d, nil
	}

	attrs := make([]sentinelAggregatorContracts.AttributeDescriptor, len(d.Attributes))
	copy(attrs, d.Attributes)
	for i := range attrs {
		resolve, ok := entry.resolveValues[attrs[i].Key]
		if !ok {
			continue
		}
		values, err := resolve(ctx, orgID)
		if err != nil {
			return sentinelAggregatorContracts.ResourceDescriptor{},
				fmt.Errorf("resolving values for %s.%s: %w", d.Type, attrs[i].Key, err)
		}
		attrs[i].Values = values
	}
	d.Attributes = attrs
	return d, nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetContextPalette implements [SentinelAggregatorServiceInterface]. The palette
// is static app vocabulary, defined next to the keys the engine stamps in
// sentinel_contexts; orgID/emplID are carried for future per-org value lists.
func (s *sentinelAggregatorService) GetContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error) {
	return sentinelContexts.BuilderPalette(), nil
}

/* ---------------------------------- <***> --------------------------------- */

// GetResourceDetails implements [SentinelAggregatorServiceInterface].
//
// One map lookup resolves both the descriptor and the lister bound to this
// type, so no module is asked to dispatch on a resource type the aggregator has
// already resolved.
func (s *sentinelAggregatorService) GetResourceDetails(
	ctx context.Context,
	orgID, emplID uuid.UUID,
	resourceType string,
	filter sentinelAggregatorContracts.InstanceFilter,
) (ResourceDetailsView, error) {
	entry, ok := s.resources[resourceType]
	if !ok {
		return ResourceDetailsView{}, fmt.Errorf("%w: %q", ErrUnknownResourceType, resourceType)
	}

	described, err := s.describe(ctx, orgID, entry)
	if err != nil {
		return ResourceDetailsView{}, err
	}

	view := ResourceDetailsView{
		ResourceDescriptor: described,
		HasInstances:       entry.list != nil,
		Instances:          []sentinelAggregatorContracts.ResourceInstance{},
	}
	if entry.list == nil {
		// No instances by design: scope through this type's container instead.
		return view, nil
	}

	page, err := entry.list(ctx, orgID, emplID, filter)
	if err != nil {
		return ResourceDetailsView{}, fmt.Errorf("listing instances of %q: %w", resourceType, err)
	}
	if page.Instances != nil {
		view.Instances = page.Instances
	}
	view.NextCursor = page.Cursor
	return view, nil
}

/* ---------------------------------- <***> --------------------------------- */

// ValidateActions implements [SentinelAggregatorServiceInterface].
//
// Called by the permissions endpoint before generating any policy text. It is
// what turns "an over-broad grant is forbidden" into "an over-broad grant is
// unrepresentable": the only action ids that can reach stored policy text are
// ones a module declared.
func (s *sentinelAggregatorService) ValidateActions(actions []string) error {
	for _, a := range actions {
		if _, ok := s.actionOwner[a]; !ok {
			return fmt.Errorf("%w: %q", ErrUnknownAction, a)
		}
	}
	return nil
}

/* ---------------------------------- <***> --------------------------------- */

// ResolveFacts implements [SentinelAggregatorServiceInterface].
//
// A type with no declared attributes has no resolver, which is not an error —
// it simply has no facts beyond its identity, and the caller falls back to the
// request's org for tenancy.
func (s *sentinelAggregatorService) ResolveFacts(
	ctx context.Context, orgID uuid.UUID, resourceType string, instanceIDs []string,
) (map[string]sentinelAggregatorContracts.ResourceFacts, error) {
	entry, ok := s.resources[resourceType]
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownResourceType, resourceType)
	}
	if entry.resolveFacts == nil || len(instanceIDs) == 0 {
		return map[string]sentinelAggregatorContracts.ResourceFacts{}, nil
	}
	facts, err := entry.resolveFacts(ctx, orgID, instanceIDs)
	if err != nil {
		return nil, fmt.Errorf("resolving facts for %q: %w", resourceType, err)
	}
	return facts, nil
}
