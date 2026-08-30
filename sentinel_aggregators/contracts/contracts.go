// Package sentinelAggregatorContracts defines the contract between the Sentinel
// resources aggregator and each module's mod_resources package.
//
// A module describes the resources it owns — Cedar entity types, the actions
// that can be performed on them, the attributes policies may condition on — and
// answers three questions about live instances: how to list them, what
// container they sit in, and what their attribute values are.
//
// # The rule this contract exists to enforce
//
// Anything DECLARED here must be ANSWERABLE here. The failure this prevents is
// specific and was live in this codebase: a module advertised `city` as a
// conditionable attribute, no call site ever passed it, and Cedar's response to
// a missing attribute is to skip the whole policy. The rule read correctly,
// compiled, stored, and displayed as active while matching nothing — with no
// error at any layer. Declaring an attribute now requires a resolver for it,
// checked at bootstrap.
//
// # Registration vs. wire shape
//
// ResourceRegistration is what a module REGISTERS; ResourceDescriptor is what
// the frontend SEES. They are separate on purpose: the registration carries
// resolver funcs, which cannot be serialized, and keeping them apart means a
// field added for the aggregator's own use never leaks into the API.
package sentinelAggregatorContracts

import (
	"context"

	sentinelAggregatorTypes "sharpnr-engine/sentinel_aggregators/types"

	"github.com/google/uuid"
)

/* -------------------------------------------------------------------------- */
/*                          What the frontend sees                            */
/* -------------------------------------------------------------------------- */

// ActionDescriptor is one action the policy builder can grant on a resource.
//
// Name is the Cedar action id embedded in policy text (`action ==
// Action::"<name>"`). Cedar action ids are a FLAT, GLOBAL namespace — there is
// no resource qualifier — so a name used by two resources means a policy scoped
// only by action reaches both. Follow can<Verb><ResourceType>; the aggregator
// rejects collisions at bootstrap either way.
//
// Names are embedded in stored policy text, so they must stay stable once
// released: renaming one silently breaks every policy that mentions it.
type ActionDescriptor struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`

	// RequiresInstance is true when the action targets a specific instance
	// (edit THIS property) rather than the type (create A property).
	//
	// It carries real weight in the generator, not just the picker: an action
	// that does not target an instance cannot be narrowed to one, so when a
	// resource is scoped to selected instances, these still emit one unscoped
	// statement. "Add a new property" has no property to pin it to.
	RequiresInstance bool `json:"requiresInstance"`
}

// AttributeValue is one value a closed-set attribute accepts.
type AttributeValue struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// AttributeDescriptor is one attribute a policy condition can test
// (`resource.<key> == ...`).
//
// This is deliberately the SAME shape the context palette uses for
// `context.*` and `principal.*` keys. There were two shapes for one widget:
// resource attributes arrived as {name, type, description} with no operators
// and no values, while context keys arrived with values but a different field
// layout — so the builder had to render the same dropdown row twice.
type AttributeDescriptor struct {
	// Key is the attribute name, matching what ResolveFacts returns.
	Key string `json:"key"`
	// Cedar is how policy text reads it, e.g. "resource.city".
	Cedar       string                             `json:"cedar"`
	Label       string                             `json:"label"`
	Description string                             `json:"description,omitempty"`
	Type        sentinelAggregatorTypes.AttrType   `json:"type"`
	Operators   []sentinelAggregatorTypes.Operator `json:"operators"`

	// Values is the closed set this attribute accepts, when there is one.
	//
	// Prefer a closed set. A free-text value is how you get a rule that reads
	// correctly and matches nothing: Cedar's `==` is case-sensitive, so an
	// admin typing "mumbai" against a stored "Mumbai" gets silence, not an
	// error.
	//
	// Populated either statically at registration or per-org by a ValueResolver
	// — the aggregator fills this in before the descriptor reaches the wire, so
	// the frontend sees one shape either way.
	Values []AttributeValue `json:"values,omitempty"`
}

// ValueResolver supplies an attribute's closed value set for one organization.
//
// For values that are data rather than constants: the cities a chain actually
// operates in, not every city in the world. Resolved when the catalog is read,
// not on the authorization path.
type ValueResolver func(ctx context.Context, orgID uuid.UUID) ([]AttributeValue, error)

// ResourceDescriptor describes one protectable resource a module exposes. Type
// is the Cedar entity type (`resource is <Type>`, `<Type>::"<id>"`). Pure data —
// this is exactly what the builder receives.
type ResourceDescriptor struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`

	// Category groups the resource in the builder ("Operations",
	// "Administration"). Without it every resource type in the product lands in
	// one flat list, which is not something a non-technical admin can scan.
	Category string `json:"category"`

	// HousedIn is the container type this resource sits in ("Property"), empty
	// for organization-level resources. Exposed to the frontend so the builder
	// can explain that scoping to a property covers everything beneath it.
	HousedIn string `json:"housedIn,omitempty"`

	Actions    []ActionDescriptor    `json:"actions"`
	Attributes []AttributeDescriptor `json:"attributes"`
}

// ResourceInstance is one live, selectable instance within an org.
type ResourceInstance struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

/* -------------------------------------------------------------------------- */
/*                        What a module must answer                           */
/* -------------------------------------------------------------------------- */

// InstanceFilter narrows an instance listing.
//
// The filter is the signature's whole reason for existing. Returning every
// instance works at twenty properties and breaks at two thousand, and a picker
// rendering a stored policy needs to resolve four specific ids without fetching
// the rest.
type InstanceFilter struct {
	// Query is free-text search over the instance label.
	Query string
	// IDs, when set, restricts the result to exactly these instances — used to
	// render a stored instance-scoped policy.
	IDs []string
	// Limit is the page size; zero means the module's default.
	Limit int
	// Cursor continues a previous page.
	Cursor string
}

// InstancePage is one page of instances.
type InstancePage struct {
	Instances []ResourceInstance
	// Cursor is the token for the next page, empty when exhausted.
	Cursor string
}

// InstanceLister returns one resource type's live instances for an org.
//
// It is bound to a single type at registration, so no module ever switches on a
// resource type: the aggregator has already resolved which lister to call.
//
// emplID is who is asking. Narrow the results to what they may see — this is
// the instance picker, so it is exactly the screen a restricted user should not
// be able to read every name out of.
type InstanceLister func(ctx context.Context, orgID, emplID uuid.UUID, filter InstanceFilter) (InstancePage, error)

// ResourceFacts is what the engine needs to know about one live instance.
type ResourceFacts struct {
	// OrgID is the organization that OWNS this instance. The engine denies any
	// request whose resource belongs to a different org before a single policy
	// runs — the tenant boundary is structural, not something a `when` clause
	// has to remember.
	OrgID uuid.UUID

	// Attributes are the values behind ResourceDescriptor.Attributes. Keys must
	// match the declared Key exactly; a mismatch means the condition reads a
	// missing attribute and Cedar skips the policy.
	Attributes map[string]any
}

// FactResolver answers, for specific instances, who owns them and what their
// declared attributes are.
//
// Required whenever the descriptor declares attributes — see the package doc.
// Batched on purpose: a list endpoint authorizing fifty rows must not fire
// fifty queries. Instances the caller cannot see, or that no longer exist, are
// simply absent from the returned map.
type FactResolver func(ctx context.Context, orgID uuid.UUID, instanceIDs []string) (map[string]ResourceFacts, error)

// ContainerResolver answers, for specific instances, which container each sits
// in — returning the container's INSTANCE ID, whose type is HousedIn.
//
// Only the direct container. The aggregator follows HousedIn up the chain and
// assembles the full ancestry itself, because it is the only component that
// sees every module.
type ContainerResolver func(ctx context.Context, orgID uuid.UUID, instanceIDs []string) (map[string]string, error)

/* -------------------------------------------------------------------------- */
/*                          What a module registers                           */
/* -------------------------------------------------------------------------- */

// ResourceRegistration is one resource type a module contributes to the
// catalog. Build one with Standard() rather than by hand — see builder.go.
type ResourceRegistration struct {
	Descriptor ResourceDescriptor

	// HousedIn is the container type, empty for organization-level resources.
	// The Organization edge is stamped by the engine from the request's org, so
	// no module ever declares it.
	HousedIn string

	// ResolveContainer is required when HousedIn is set.
	ResolveContainer ContainerResolver

	// ResolveFacts is required when Descriptor.Attributes is non-empty.
	ResolveFacts FactResolver

	// ResolveValues supplies per-org value lists, keyed by attribute Key. An
	// attribute absent from this map keeps whatever static Values it declared.
	ResolveValues map[string]ValueResolver

	// ListInstances is optional. Nil means the type has no selectable instances
	// — policies about it are type-level, or it is reached by scoping to its
	// container. That is different from a lister returning an empty page, which
	// means "instances exist, this org has none".
	//
	// Most types should leave this nil. Nobody scopes a policy to one menu
	// item; they scope to the property and let containment do the rest.
	ListInstances InstanceLister
}

// ResourceProvider is what a module's mod_resources package implements and
// registers with the aggregator at bootstrap. One provider may contribute
// several resource types.
type ResourceProvider interface {
	Resources() []ResourceRegistration
}
