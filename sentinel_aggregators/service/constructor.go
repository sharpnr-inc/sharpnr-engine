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
// into one catalog: what the policy builder can offer, and what the Guard needs
// to assemble a complete authorization request. It makes no authorization
// decisions itself.
type SentinelAggregatorServiceInterface interface {
	// ListResources returns a summary of every registered resource type.
	ListResources(ctx context.Context, orgID, emplID uuid.UUID) ([]ResourceSummaryView, error)

	// GetCatalog returns every resource WITH its actions and attributes, in one
	// response. This is what the permission matrix reads — ListResources omits
	// actions, so without this the frontend fans out one request per type.
	GetCatalog(ctx context.Context, orgID, emplID uuid.UUID) ([]sentinelAggregatorContracts.ResourceDescriptor, error)

	// GetResourceDetails returns one resource's full palette plus a page of its
	// live instances.
	GetResourceDetails(ctx context.Context, orgID, emplID uuid.UUID, resourceType string, filter sentinelAggregatorContracts.InstanceFilter) (ResourceDetailsView, error)

	// GetContextPalette returns the context vocabulary restrictions can be
	// authored against.
	GetContextPalette(ctx context.Context, orgID, emplID uuid.UUID) (sentinelContexts.Palette, error)

	// ValidateActions reports an error unless every id is a registered action.
	// This is what makes an over-broad grant unrepresentable rather than merely
	// forbidden: the permissions endpoint cannot persist an action nobody
	// declared.
	ValidateActions(actions []string) error

	// ResolveFacts returns the owning org and declared attributes for instances
	// of one type, for the Guard to stamp onto the request.
	ResolveFacts(ctx context.Context, orgID uuid.UUID, resourceType string, instanceIDs []string) (map[string]sentinelAggregatorContracts.ResourceFacts, error)

	// ResolveAncestry returns each instance's full container chain, nearest
	// first, by walking HousedIn across modules.
	ResolveAncestry(ctx context.Context, orgID uuid.UUID, resourceType string, instanceIDs []string) (map[string][]ResourceRef, error)
}

// ResourceRef names one instance of one resource type.
type ResourceRef struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// registeredResource is one catalog entry.
type registeredResource struct {
	descriptor       sentinelAggregatorContracts.ResourceDescriptor
	housedIn         string
	resolveContainer sentinelAggregatorContracts.ContainerResolver
	resolveFacts     sentinelAggregatorContracts.FactResolver
	resolveValues    map[string]sentinelAggregatorContracts.ValueResolver
	list             sentinelAggregatorContracts.InstanceLister
}

type sentinelAggregatorService struct {
	// order preserves registration order so the catalog is stable across calls.
	order     []string
	resources map[string]registeredResource
	// actionOwner maps every registered action id to the resource that owns it,
	// which doubles as the ValidateActions lookup.
	actionOwner map[string]string
}

var _ SentinelAggregatorServiceInterface = (*sentinelAggregatorService)(nil)

// grandfatheredActions predate can<Verb><ResourceType> and are already embedded
// in stored policy text. Renaming one would silently break every policy that
// mentions it, so they are exempt from the naming check rather than fixed.
var grandfatheredActions = map[string]bool{
	"canAddBusinessModules": true,
	"canAddInfraModules":    true,
}

// NewSentinelAggregatorService builds the catalog from the given providers,
// validating it as it goes.
//
// Every problem below is a WIRING BUG that would otherwise surface as a broken
// policy builder or, worse, a rule that silently matches nothing while
// displaying as active. The catalog is assembled once at bootstrap and never
// changes, so this panics rather than serving something wrong: a failed boot is
// loud, a silently dead permission rule is not.
func NewSentinelAggregatorService(providers ...sentinelAggregatorContracts.ResourceProvider) *sentinelAggregatorService {
	s := &sentinelAggregatorService{
		resources:   make(map[string]registeredResource),
		actionOwner: make(map[string]string),
	}

	for _, provider := range providers {
		for _, reg := range provider.Resources() {
			s.register(reg)
		}
	}

	// Deferred to the end because a module may be housed in a type registered
	// by a provider that comes later in the list.
	s.validateHierarchy()
	return s
}

/* ---------------------------------- <***> --------------------------------- */

func (s *sentinelAggregatorService) register(reg sentinelAggregatorContracts.ResourceRegistration) {
	d := reg.Descriptor
	resourceType := strings.TrimSpace(d.Type)

	if resourceType == "" {
		panic("sentinel aggregator: a resource was registered with an empty Type")
	}
	if _, dup := s.resources[resourceType]; dup {
		panic(fmt.Sprintf("sentinel aggregator: resource type %q registered twice", resourceType))
	}
	if strings.TrimSpace(d.Label) == "" {
		panic(fmt.Sprintf("sentinel aggregator: resource %q has no Label; it is what an administrator reads", resourceType))
	}
	if strings.TrimSpace(d.Category) == "" {
		panic(fmt.Sprintf(
			"sentinel aggregator: resource %q has no Category. Without one every resource in the "+
				"product lands in a single flat list.", resourceType))
	}

	s.validateActions(resourceType, d.Actions)
	validateAttributes(resourceType, d.Attributes)

	// Declared must mean answerable — the rule this contract exists for.
	if len(d.Attributes) > 0 && reg.ResolveFacts == nil {
		panic(fmt.Sprintf(
			"sentinel aggregator: resource %q declares %d attribute(s) but no ResolveFacts. "+
				"Cedar skips a whole policy when an attribute is missing, so every condition "+
				"using them would read as active and match nothing.",
			resourceType, len(d.Attributes)))
	}
	declared := make(map[string]struct{}, len(d.Attributes))
	for _, a := range d.Attributes {
		declared[a.Key] = struct{}{}
	}
	for key := range reg.ResolveValues {
		if _, ok := declared[key]; !ok {
			panic(fmt.Sprintf(
				"sentinel aggregator: resource %q has a value resolver for %q, which it does not declare as an attribute",
				resourceType, key))
		}
	}
	if reg.HousedIn != "" && reg.ResolveContainer == nil {
		panic(fmt.Sprintf(
			"sentinel aggregator: resource %q is HousedIn %q but has no ResolveContainer, so its "+
				"instances can never be placed in a container.", resourceType, reg.HousedIn))
	}

	s.resources[resourceType] = registeredResource{
		descriptor:       d,
		housedIn:         reg.HousedIn,
		resolveContainer: reg.ResolveContainer,
		resolveFacts:     reg.ResolveFacts,
		resolveValues:    reg.ResolveValues,
		list:             reg.ListInstances,
	}
	s.order = append(s.order, resourceType)
}

/* ---------------------------------- <***> --------------------------------- */

// validateActions rejects empty, duplicated and cross-resource-colliding ids.
//
// The cross-resource check is the important one. Cedar has no resource
// qualifier on an action: `permit(principal, action == Action::"canView",
// resource)` matches that id on EVERY resource type. If two modules both
// register "canView", a policy meant for one silently reaches the other. The
// aggregator is the only place that sees all modules at once, so it is the only
// place this can be caught.
func (s *sentinelAggregatorService) validateActions(resourceType string, actions []sentinelAggregatorContracts.ActionDescriptor) {
	if len(actions) == 0 {
		panic(fmt.Sprintf("sentinel aggregator: resource %q registered no actions", resourceType))
	}
	for _, a := range actions {
		name := strings.TrimSpace(a.Name)
		if name == "" {
			panic(fmt.Sprintf("sentinel aggregator: resource %q has an action with an empty Name", resourceType))
		}
		if strings.TrimSpace(a.Label) == "" {
			panic(fmt.Sprintf("sentinel aggregator: action %q on %q has no Label", name, resourceType))
		}
		if owner, taken := s.actionOwner[name]; taken {
			if owner == resourceType {
				panic(fmt.Sprintf("sentinel aggregator: resource %q declares action %q twice", resourceType, name))
			}
			panic(fmt.Sprintf(
				"sentinel aggregator: action %q is registered by both %q and %q. Cedar action ids are "+
					"global, so a policy scoped to this action would reach both resources. Use "+
					"can<Verb><ResourceType>.",
				name, owner, resourceType))
		}
		// The type name must appear in the id — Contains rather than HasSuffix so
		// that a trailing qualifier still reads naturally
		// (canManageSentinelGroupMembers, canEditPropertyDetails). Embedding the
		// type is what keeps ids unique in Cedar's flat action namespace; where
		// exactly it sits does not matter.
		if !grandfatheredActions[name] && !strings.Contains(name, resourceType) {
			panic(fmt.Sprintf(
				"sentinel aggregator: action %q on %q does not contain the resource type. Follow "+
					"can<Verb><ResourceType>; it is what keeps ids unique in Cedar's flat action "+
					"namespace. If this id is already in stored policy text, add it to "+
					"grandfatheredActions instead of renaming it.",
				name, resourceType))
		}
		s.actionOwner[name] = resourceType
	}
}

// validateAttributes rejects empty names, duplicates within a resource, and
// types or operators the policy builder cannot render.
func validateAttributes(resourceType string, attributes []sentinelAggregatorContracts.AttributeDescriptor) {
	seen := make(map[string]struct{}, len(attributes))
	for _, attr := range attributes {
		key := strings.TrimSpace(attr.Key)
		if key == "" {
			panic(fmt.Sprintf("sentinel aggregator: resource %q has an attribute with an empty Key", resourceType))
		}
		if _, dup := seen[key]; dup {
			panic(fmt.Sprintf("sentinel aggregator: resource %q declares attribute %q twice", resourceType, key))
		}
		if !attr.Type.Valid() {
			panic(fmt.Sprintf(
				"sentinel aggregator: resource %q attribute %q has type %q; valid types are %v",
				resourceType, key, attr.Type, sentinelAggregatorTypes.AllAttrTypes))
		}
		if len(attr.Operators) == 0 {
			panic(fmt.Sprintf("sentinel aggregator: resource %q attribute %q offers no operators", resourceType, key))
		}
		for _, op := range attr.Operators {
			if !op.Valid() {
				panic(fmt.Sprintf("sentinel aggregator: resource %q attribute %q has invalid operator %q", resourceType, key, op))
			}
		}
		seen[key] = struct{}{}
	}
}

/* ---------------------------------- <***> --------------------------------- */

// validateHierarchy checks that every HousedIn names a registered type and that
// the containment graph has no cycle.
//
// A cycle would make ResolveAncestry loop forever on a live request, so it is
// worth a few lines at boot rather than a depth cap in the hot path.
func (s *sentinelAggregatorService) validateHierarchy() {
	for _, resourceType := range s.order {
		parent := s.resources[resourceType].housedIn
		if parent == "" {
			continue
		}
		if _, ok := s.resources[parent]; !ok {
			panic(fmt.Sprintf(
				"sentinel aggregator: resource %q is HousedIn %q, which no module registered",
				resourceType, parent))
		}

		seen := map[string]bool{resourceType: true}
		for cur := parent; cur != ""; cur = s.resources[cur].housedIn {
			if seen[cur] {
				panic(fmt.Sprintf("sentinel aggregator: containment cycle through %q", cur))
			}
			seen[cur] = true
		}
	}
}
