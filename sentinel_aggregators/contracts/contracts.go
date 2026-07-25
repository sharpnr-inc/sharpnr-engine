// Package sentinelAggregatorContracts defines the contract between the Sentinel
// resources aggregator and each module's mod_resources package. A module
// describes the resources it owns (Cedar entity types, the actions that can be
// performed on them, the attributes policies may condition on) and can list the
// live instances of those resources for an org — everything the frontend policy
// builder needs to author a Cedar policy.
package sentinelAggregatorContracts

import (
	"context"

	"github.com/google/uuid"
)

// ActionDescriptor is one action the policy builder can grant or forbid on a
// resource. Name is the Cedar action id embedded in policy text
// (`action == Action::"<name>"`); it must stay stable once released.
type ActionDescriptor struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	// RequiresInstance is true when the action targets a specific resource
	// instance (edit/view one property) rather than the type (create/list).
	RequiresInstance bool `json:"requiresInstance"`
}

// AttributeDescriptor is one resource attribute a policy condition can test
// (`resource.<name> == ...`). Type is the Cedar primitive: "string", "long" or
// "bool".
type AttributeDescriptor struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ResourceDescriptor describes one protectable resource a module exposes. Type
// is the Cedar entity type (`resource is <Type>`, `<Type>::"<id>"`).
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

// ResourceProvider is what a module's mod_resources package implements and
// registers with the aggregator at bootstrap. One provider may expose several
// resource types (the Sentinel module itself exposes policy, role and group).
type ResourceProvider interface {
	// Resources returns the static descriptors of every resource type the
	// module owns.
	Resources() []ResourceDescriptor

	// ListInstances returns the org's live instances of one of the provider's
	// resource types. Providers with no meaningful instances return an empty
	// slice.
	ListInstances(ctx context.Context, orgID, emplID uuid.UUID, resourceType string) ([]ResourceInstance, error)
}
