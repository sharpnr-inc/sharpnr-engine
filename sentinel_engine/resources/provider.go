// Package sentinel_resources is the resource-centric heart of the policy builder.
// Each feature module that owns a resource type registers two things here:
//
//   - a Descriptor: the static metadata for that resource type — its actions and
//     its ABAC attributes — which the FE renders when an admin picks a resource.
//   - a Provider: the dynamic lister of the org's live instances of that type
//     (e.g. each Property::"<id>"), which the FE offers for `resource in ...`.
//
// Actions live under the resource they apply to (no separate flat catalog); the
// engine's flat action-validity set is fed from every descriptor.
package sentinel_resources

import (
	"context"
	"fmt"
	"sort"
	"sync"

	sentinel_actions "sharpnr-engine/sentinel_engine/actions"

	"github.com/google/uuid"
)

// Attribute describes one ABAC attribute of a resource, so the builder can offer
// `when { resource.<Name> ... }`. Type is the Cedar value type the FE renders an
// input for: "String" | "Long" | "Boolean" | "Set" | "EntityRef".
type Attribute struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Action is one action applicable to a resource type. Key is the Cedar action id
// baked into policy text; renaming it breaks stored policies that reference it.
type Action struct {
	Key         sentinel_actions.Name `json:"key"`
	Label       string                `json:"label"`
	Description string                `json:"description"`
}

// Descriptor is a resource type's full policy-builder metadata. Type is the Cedar
// entity type (must match what the enforcement call site passes as Resource.Type
// and what stored policies reference, e.g. "Property").
type Descriptor struct {
	Type        string      `json:"type"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
	Actions     []Action    `json:"actions"`
	Attributes  []Attribute `json:"attributes"`
}

// Ref is one selectable resource instance for the builder's `resource in
// <Type>::"<id>"` picker. Type is the Cedar entity type, ID the entity id, Label
// the human name.
type Ref struct {
	ID    uuid.UUID `json:"id"`
	Type  string    `json:"type"`
	Label string    `json:"label"`
}

// Provider lists the live instances of one resource type within an org. Each
// owning module implements it (e.g. the property module over org_property) and
// registers it at startup.
type Provider interface {
	// Type is the Cedar entity type this provider serves, e.g. "Property".
	Type() string
	// List returns the org's instances of that type, scoped to the tenant.
	List(ctx context.Context, orgID uuid.UUID) ([]Ref, error)
}

var (
	mu          sync.RWMutex
	descriptors = map[string]Descriptor{}
	providers   = map[string]Provider{}
)

// RegisterDescriptor installs a resource type's static metadata and feeds its
// action ids into the engine's action-validity set. Last registration wins. Call
// at startup (composition root), before the builder is served.
func RegisterDescriptor(d Descriptor) {
	mu.Lock()
	descriptors[d.Type] = d
	mu.Unlock()

	keys := make([]sentinel_actions.Name, 0, len(d.Actions))
	for _, a := range d.Actions {
		keys = append(keys, a.Key)
	}
	sentinel_actions.RegisterValid(keys...)
}

// RegisterProvider installs the instance lister for a resource type. Last
// registration wins (a test can override). Call at startup.
func RegisterProvider(p Provider) {
	mu.Lock()
	defer mu.Unlock()
	providers[p.Type()] = p
}

// Descriptors returns every registered resource descriptor, sorted by Type. The
// FE lists these as the resource picker's options.
func Descriptors() []Descriptor {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]Descriptor, 0, len(descriptors))
	for _, d := range descriptors {
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Type < out[j].Type })
	return out
}

// DescriptorFor returns the descriptor for a resource type. ok is false for an
// unregistered type, so a bad ?type= surfaces instead of an empty payload.
func DescriptorFor(resourceType string) (Descriptor, bool) {
	mu.RLock()
	defer mu.RUnlock()
	d, ok := descriptors[resourceType]
	return d, ok
}

// Types returns the registered resource types, sorted.
func Types() []string {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]string, 0, len(descriptors))
	for t := range descriptors {
		out = append(out, t)
	}
	sort.Strings(out)
	return out
}

// List dispatches to the provider for resourceType and returns the org's live
// instances. Errors if no provider is registered for the type.
func List(ctx context.Context, resourceType string, orgID uuid.UUID) ([]Ref, error) {
	mu.RLock()
	p, ok := providers[resourceType]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("no resource provider registered for type %q", resourceType)
	}
	return p.List(ctx, orgID)
}
