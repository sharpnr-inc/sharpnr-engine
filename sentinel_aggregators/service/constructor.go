package sentinelAggregatorService

import (
	"context"
	"fmt"
	"strings"

	sentinelAggregatorContracts "sharpnr-engine/sentinel_aggregators/contracts"
	sentinelAggregatorTypes "sharpnr-engine/sentinel_aggregators/types"
	sentinelContexts "sharpnr-engine/sentinel_contexts"

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

	// GetContextPalette returns the context vocabulary the builder can author
	// conditions against: the canonical `context.*` request keys and the
	// `principal.*` user keys the engine stamps on every evaluation.
	GetContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error)
}

// registeredResource is one catalog entry: the descriptor the frontend sees and
// the lister bound to that type.
type registeredResource struct {
	descriptor sentinelAggregatorContracts.ResourceDescriptor
	list       sentinelAggregatorContracts.InstanceLister
}

type sentinelAggregatorService struct {
	// order preserves registration order so the FE sees a stable catalog.
	order     []string
	resources map[string]registeredResource
}

// Compile-time guard: the concrete service must satisfy the interface.
var _ SentinelAggregatorServiceInterface = (*sentinelAggregatorService)(nil)

// NewSentinelAggregatorService builds the catalog from the given providers,
// validating it as it goes.
//
// Every problem below is a WIRING BUG that would otherwise surface as a broken
// policy builder or, worse, a policy that grants more than its author intended.
// The catalog is assembled once at bootstrap and never changes, so this panics
// rather than serving something wrong: a failed boot is loud, a silently
// over-broad action is not.
func NewSentinelAggregatorService(providers ...sentinelAggregatorContracts.ResourceProvider) *sentinelAggregatorService {
	s := &sentinelAggregatorService{
		resources: make(map[string]registeredResource),
	}

	// Cedar action ids are a flat global namespace, so this map spans every
	// resource of every module — not just the one being registered.
	actionOwner := make(map[string]string)

	for _, provider := range providers {
		for _, reg := range provider.Resources() {
			d := reg.Descriptor
			resourceType := strings.TrimSpace(d.Type)

			if resourceType == "" {
				panic("sentinel aggregator: a resource was registered with an empty Type")
			}
			if _, dup := s.resources[resourceType]; dup {
				panic(fmt.Sprintf("sentinel aggregator: resource type %q registered twice", resourceType))
			}
			validateActions(resourceType, d.Actions, actionOwner)
			validateAttributes(resourceType, d.Attributes)

			s.resources[resourceType] = registeredResource{descriptor: d, list: reg.ListInstances}
			s.order = append(s.order, resourceType)
		}
	}
	return s
}

/* ---------------------------------- <***> --------------------------------- */

// validateActions rejects empty or duplicated action ids.
//
// The cross-resource check is the important one. Cedar has no resource
// qualifier on an action: `permit(principal, action == Action::"canView",
// resource)` matches that id on EVERY resource type. If two modules both
// register "canView", a policy meant for one of them silently reaches the
// other. The aggregator is the only place that sees all modules at once, so it
// is the only place this can be caught.
func validateActions(resourceType string, actions []sentinelAggregatorContracts.ActionDescriptor, actionOwner map[string]string) {
	for _, a := range actions {
		name := strings.TrimSpace(a.Name)
		if name == "" {
			panic(fmt.Sprintf("sentinel aggregator: resource %q has an action with an empty Name", resourceType))
		}
		if owner, taken := actionOwner[name]; taken {
			if owner == resourceType {
				panic(fmt.Sprintf(
					"sentinel aggregator: resource %q declares action %q twice",
					resourceType, name))
			}
			panic(fmt.Sprintf(
				"sentinel aggregator: action %q is registered by both %q and %q. "+
					"Cedar action ids are global, so a policy scoped to this action would reach both "+
					"resources. Give each a resource-suffixed name (e.g. %sOn%s).",
				name, owner, resourceType, name, resourceType))
		}
		actionOwner[name] = resourceType
	}
}

// validateAttributes rejects empty names, duplicates within a resource, and
// types the policy builder cannot render.
func validateAttributes(resourceType string, attributes []sentinelAggregatorContracts.AttributeDescriptor) {
	seen := make(map[string]struct{}, len(attributes))
	for _, attr := range attributes {
		name := strings.TrimSpace(attr.Name)
		if name == "" {
			panic(fmt.Sprintf("sentinel aggregator: resource %q has an attribute with an empty Name", resourceType))
		}
		if _, dup := seen[name]; dup {
			panic(fmt.Sprintf("sentinel aggregator: resource %q declares attribute %q twice", resourceType, name))
		}
		if !attr.Type.Valid() {
			panic(fmt.Sprintf(
				"sentinel aggregator: resource %q attribute %q has type %q; valid types are %v",
				resourceType, name, attr.Type, sentinelAggregatorTypes.AllAttrTypes))
		}
		seen[name] = struct{}{}
	}
}
