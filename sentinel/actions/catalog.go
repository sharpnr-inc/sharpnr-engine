package sentinel_actions

import "sync"

// Action validation registry.
//
// Actions are no longer a flat catalog — they are declared per resource type via
// sentinel_resources.Descriptor (the resource-centric policy builder). This file
// keeps only the engine's side of the deal: the set of action ids known to be
// valid, so a stored policy can be checked against actions that actually exist.
//
// sentinel_resources.RegisterDescriptor feeds every descriptor's action keys in
// here (one-way dependency: resources -> actions), giving the engine a single
// flat validity set without re-introducing a catalog.

var (
	registryMu sync.RWMutex
	valid      = map[Name]bool{}
)

// RegisterValid marks action ids as known/valid. Idempotent. Called by
// sentinel_resources.RegisterDescriptor at startup, before policies are served.
func RegisterValid(names ...Name) {
	registryMu.Lock()
	defer registryMu.Unlock()
	for _, n := range names {
		valid[n] = true
	}
}

// IsValid reports whether key is a registered action id. Use it to validate a
// policy's action references on write, so we never store Cedar that points at an
// action that doesn't exist.
func IsValid(key Name) bool {
	registryMu.RLock()
	defer registryMu.RUnlock()
	return valid[key]
}
