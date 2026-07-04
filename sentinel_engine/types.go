package sentinelEngine

import "github.com/google/uuid"

// Resource identifies what an action is attempted on. Type is the Cedar entity
// type a module registered through its mod_resources (e.g. "Property"); ID is
// the instance id (usually a UUID string). Leave ID empty for type-level checks
// such as create/list — the engine normalizes it to "*" so policies written as
// `resource is Property` still match.
type Resource struct {
	Type string
	ID   string
}

// AccessRequest is one authorization question: may UserID perform Action on
// Resource inside OrgID?
//
// ResourceAttributes become the resource entity's Cedar attributes so policies
// can condition on them (e.g. resource.city == "Mumbai"); Context becomes the
// request context (`context.<key>` in policy text). Both are optional.
type AccessRequest struct {
	OrgID  uuid.UUID
	UserID uuid.UUID // employee id — the Cedar principal User::"<id>"
	Action string    // Cedar action id, e.g. "canViewProperty"

	Resource           Resource
	ResourceAttributes map[string]any
	Context            map[string]any
}

// PolicyRecord is the minimal policy row the engine compiles: the stored Cedar
// text plus the row id, which doubles as the cedar.PolicyID so authorization
// diagnostics can be traced back to the DB row.
type PolicyRecord struct {
	ID   uuid.UUID
	Text string
}
