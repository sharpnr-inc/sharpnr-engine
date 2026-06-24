// Package sentinel_resources is the dynamic counterpart to the static action
// catalog. Actions are a fixed, code-declared set; resources are live data
// (an org's properties, documents, …) that must be fetched per request. A
// feature module that owns a resource type registers a Provider here; the
// sentinel layer dispatches to it so the FE can list the instances a policy may
// be scoped to (e.g. resource in [Property::"<id>"]).
package sentinel_resources

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/google/uuid"
)

// Ref is one selectable resource instance for the policy builder. Type is the
// Cedar entity type (must match what the enforcement call site passes as
// Resource.Type and what the generator emits), ID is the entity id, and Label is
// the human name the picker shows.
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
	mu        sync.RWMutex
	providers = map[string]Provider{}
)

// Register installs a provider for its resource type. Last registration wins,
// so a test can override a real provider. Call at startup (composition root).
func Register(p Provider) {
	mu.Lock()
	defer mu.Unlock()
	providers[p.Type()] = p
}

// List dispatches to the provider for resourceType and returns the org's
// instances. Errors if no provider is registered for the type, so a typo in the
// FE's ?type= surfaces instead of silently returning an empty list.
func List(ctx context.Context, resourceType string, orgID uuid.UUID) ([]Ref, error) {
	mu.RLock()
	p, ok := providers[resourceType]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("no resource provider registered for type %q", resourceType)
	}
	return p.List(ctx, orgID)
}

// Types returns the registered resource types, sorted. The FE uses it to know
// which resource kinds a policy can be scoped to.
func Types() []string {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]string, 0, len(providers))
	for t := range providers {
		out = append(out, t)
	}
	sort.Strings(out)
	return out
}
