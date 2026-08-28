// Package sentinelAggregatorContracts defines the contract between the Sentinel
// resources aggregator and each module's mod_resources package.
//
// A module describes the resources it owns — Cedar entity types, the actions
// that can be performed on them, the attributes policies may condition on — and
// says how to list the live instances of each. Together that is everything the
// frontend policy builder needs to author a Cedar policy.
//
// # Registration vs. wire shape
//
// ResourceRegistration is what a module REGISTERS; ResourceDescriptor is what
// the frontend SEES. They are separate on purpose: the registration carries an
// InstanceLister, which cannot be serialized, and keeping them apart means a
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

// ActionDescriptor is one action the policy builder can grant or forbid on a
// resource.
//
// Name is the Cedar action id embedded in policy text (`action ==
// Action::"<name>"`). Cedar action ids are a FLAT, GLOBAL namespace — there is
// no resource qualifier — so a name used by two resources means a policy scoped
// only by action reaches both. The aggregator rejects duplicates at bootstrap;
// keep names resource-suffixed (canViewProperty, not canView).
//
// Names are embedded in stored policy text, so they must stay stable once
// released: renaming one silently breaks every policy that mentions it.
type ActionDescriptor struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	// RequiresInstance is true when the action targets a specific resource
	// instance (edit/view one property) rather than the type (create/list).
	// It tells the builder whether to show an instance picker.
	RequiresInstance bool `json:"requiresInstance"`
}

// AttributeDescriptor is one resource attribute a policy condition can test
// (`resource.<name> == ...`).
//
// The name must match what the engine caller passes in
// AccessRequest.ResourceAttributes. Cedar errors on a missing attribute and
// skips the whole policy, so an attribute advertised here but never stamped at
// the call site makes every condition using it silently fail.
type AttributeDescriptor struct {
	Name        string                           `json:"name"`
	Type        sentinelAggregatorTypes.AttrType `json:"type"`
	Description string                           `json:"description"`
}

// ResourceDescriptor describes one protectable resource a module exposes. Type
// is the Cedar entity type (`resource is <Type>`, `<Type>::"<id>"`). Pure data —
// this is exactly what the builder receives.
type ResourceDescriptor struct {
	Type        string                `json:"type"`
	Label       string                `json:"label"`
	Description string                `json:"description"`
	Actions     []ActionDescriptor    `json:"actions"`
	Attributes  []AttributeDescriptor `json:"attributes"`
}

// ResourceInstance is one live, selectable instance of a resource within an org
// (e.g. a property), for policies scoped to `resource == <Type>::"<id>"`.
type ResourceInstance struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

/* -------------------------------------------------------------------------- */
/*                          What a module registers                           */
/* -------------------------------------------------------------------------- */

// InstanceLister returns one resource type's live instances for an org.
//
// It is bound to a single type at registration, so no module ever switches on a
// resource type: the aggregator has already resolved which lister to call.
//
// emplID is carried so listing can be narrowed to what the caller may see once
// the authorization guard is wired; it is unused today.
type InstanceLister func(ctx context.Context, orgID, emplID uuid.UUID) ([]ResourceInstance, error)

// ResourceRegistration is one resource type a module contributes to the
// catalog: what it looks like, and how to list its instances.
//
// A nil ListInstances means the type has no selectable instances — policies
// about it are type-level only. That is different from a lister returning an
// empty slice, which means "instances exist, this org has none".
type ResourceRegistration struct {
	Descriptor    ResourceDescriptor
	ListInstances InstanceLister
}

// ResourceProvider is what a module's mod_resources package implements and
// registers with the aggregator at bootstrap. One provider may contribute
// several resource types (the Sentinel module registers policy, role and group).
type ResourceProvider interface {
	Resources() []ResourceRegistration
}
