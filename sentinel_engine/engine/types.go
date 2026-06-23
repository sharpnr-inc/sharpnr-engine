package sentinelengine

import (
	"github.com/cedar-policy/cedar-go/types"
	"github.com/google/uuid"
)

// Resource identifies the thing being acted upon.
type Resource struct {
	Type    string                 // Cedar entity type, e.g. "Document"
	ID      string                 // entity id
	Attrs   map[string]types.Value // optional attributes; orgId is added automatically
	Parents []types.EntityUID      // optional hierarchy, e.g. a Document in a Folder
}

// AuthorizeInput is everything needed to make one decision.
type AuthorizeInput struct {
	OrgID    uuid.UUID
	UserID   uuid.UUID
	Action   string // a sentinel_actions constant, e.g. sentinel_actions.ActionEditDocument
	Resource Resource
	Context  map[string]types.Value // optional request context
}

// Result is the outcome of a check.
type Result struct {
	Allowed bool
	Reasons []string // ids of the policies that determined the decision
	Errors  []string // evaluation errors, if any (the decision is still authoritative)
}
