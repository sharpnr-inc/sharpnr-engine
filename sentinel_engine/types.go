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

	// OrgID is the organization that OWNS this resource instance, set when the
	// caller has loaded the row and knows it. The engine denies any request
	// whose resource belongs to a different org than the request's org — the
	// multi-tenant boundary is enforced structurally, before any policy runs,
	// so no policy text (and no forgotten `when` clause) can cross it. Callers
	// that resolve instances by user-supplied ids should set this.
	OrgID uuid.UUID
}

// AccessRequest is one authorization question: may UserID perform Action on
// Resource inside OrgID?
//
// ResourceAttributes become the resource entity's Cedar attributes so policies
// can condition on them (e.g. resource.city == "Mumbai"); Context becomes the
// request context (`context.<key>` in policy text). Both are optional.
//
// The engine stamps tenancy attributes itself: principal.orgId is always the
// request's org, and resource.orgId is Resource.OrgID when set (the request's
// org otherwise). Authored policies may therefore condition on
// `resource.orgId == principal.orgId`, though the engine already enforces that
// boundary structurally.
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
