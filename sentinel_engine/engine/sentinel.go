// Package sentinelengine is the authorization engine. It uses cedar-go (v1.8.0)
// and follows the model at https://docs.cedarpolicy.com/.
//
// Two things move at different speeds, so they're handled differently:
//   - Policies change rarely -> compiled into a per-org *cedar.PolicySet and
//     cached (see cache.go).
//   - Membership (who is in which role/group) and principal attributes change
//     constantly -> the EntityMap is rebuilt from the database on every
//     authorization call (see entities.go).
//
// The package is split for readability:
//   - sentinel.go    — the Sentinel type, its constructor, and entity-type names
//   - repository.go  — the Repository interface the engine reads through
//   - types.go       — the request/result DTOs (Resource, AuthorizeInput, Result)
//   - authorize.go   — IsAuthorized: the single-decision entry point
//   - cache.go       — the per-org compiled PolicySet cache + invalidation
//   - entities.go    — buildEntities: the principal/resource entity graph
package sentinelengine

import (
	"sync"

	"github.com/cedar-policy/cedar-go"
	"github.com/google/uuid"
)

// Cedar entity type names. These must match the types referenced in stored
// policies (e.g. `principal in Role::"owner"`).
const (
	entityOrganization = "Organization"
	entityUser         = "User"
	entityRole         = "Role"
	entityGroup        = "Group"
)

// Sentinel is the engine. Safe for concurrent use.
type Sentinel struct {
	repo Repository

	mu    sync.RWMutex
	cache map[uuid.UUID]*cedar.PolicySet // compiled policies, keyed by org
}

// New constructs a Sentinel engine.
func New(repo Repository) *Sentinel {
	return &Sentinel{
		repo:  repo,
		cache: make(map[uuid.UUID]*cedar.PolicySet),
	}
}
