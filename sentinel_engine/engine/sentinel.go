package sentinelengine

import (
	"context"
	"fmt"
	"strings"

	sentinelengine_models "sharpnr-engine/sentinel_engine/models"

	cedar "github.com/cedar-policy/cedar-go"
)

// Sentinel is SharpNR's authorization engine. Give it a question —
// "can THIS principal do THIS action on THIS resource?" — and it answers
// Allowed or Denied using Cedar.
//
// It does NOT talk to your database directly. Instead it asks a PolicyStore
// (which you implement against your repo) for the two things Cedar needs:
//
//  1. the policy rules that apply to the principal, and
//  2. the roles/groups the principal belongs to (so membership-based rules match).
//
// This keeps the Cedar mechanics in one place and your data access in another.

/* -------------------------------------------------------------------------- */
/*                              public request/result                         */
/* -------------------------------------------------------------------------- */

// EvaluationRequest is one authorization question.
type EvaluationRequest struct {
	// Who is asking. PrincipalType is almost always "User".
	PrincipalType string
	PrincipalID   string

	// What they want to do, e.g. "read", "delete".
	Action string

	// What they want to do it to, e.g. ResourceType "Restaurant", ResourceID "<uuid>".
	ResourceType string
	ResourceID   string

	// Optional parent of the resource, e.g. the Property a Restaurant belongs to:
	// ResourceParentType "Property", ResourceParentID "<uuid>". This is what lets
	// a per-property policy ("...resource is Restaurant in Property::X") match.
	// The calling service supplies it because it knows where the resource lives.
	ResourceParentType string
	ResourceParentID   string

	// Context is optional extra data for attribute-based rules (ABAC). It maps
	// to Cedar's `context` record, so a policy can say `when { context.mfa == true }`.
	// Only string, bool, and int values are supported here.
	Context map[string]any
}

// Decision is the engine's answer. Allowed is the only field business code
// needs; Reason carries a short diagnostic for logging/debugging.
type Decision struct {
	Allowed bool
	Reason  string
}

/* -------------------------------------------------------------------------- */
/*                                 policy store                               */
/* -------------------------------------------------------------------------- */

// PolicyStore is how the engine reaches your data. Implement it against your
// repo (sentinel_engine/repo). The engine never assumes how policies are
// stored — only that you can answer these two questions for a principal.
type PolicyStore interface {
	// PoliciesFor returns every policy that applies to this principal, already
	// merged from the principal's roles and groups. Return them as models.Policy;
	// the engine converts each to Cedar text via ToCedar().
	PoliciesFor(ctx context.Context, principalType, principalID string) ([]sentinelengine_models.Policy, error)

	// MembershipsOf returns the roles and groups this principal belongs to, as
	// (type, id) pairs — e.g. {"Role","engineer"}, {"Group","billing-team"}.
	// The engine records these as Cedar "parents" so that a rule written for a
	// Role/Group matches a User who is a member.
	MembershipsOf(ctx context.Context, principalType, principalID string) ([]Membership, error)
}

// Membership is one role/group a principal belongs to.
type Membership struct {
	Type string // "Role" or "Group"
	ID   string
}

/* -------------------------------------------------------------------------- */
/*                                 the engine                                 */
/* -------------------------------------------------------------------------- */

type Sentinel struct {
	store PolicyStore
}

// New builds an engine backed by the given store.
func New(store PolicyStore) *Sentinel {
	return &Sentinel{store: store}
}

// Boot is the one-call startup helper referenced in the docs. Wire your
// repo-backed PolicyStore here and inject the returned *Sentinel into services.
func Boot(store PolicyStore) *Sentinel {
	return New(store)
}

// Evaluate answers a single authorization question. The five steps below are
// the whole flow; everything else in this file is plumbing for them.
func (s *Sentinel) Evaluate(ctx context.Context, req EvaluationRequest) (Decision, error) {
	// 1. Ask the store for the policies and memberships of this principal.
	policies, err := s.store.PoliciesFor(ctx, req.PrincipalType, req.PrincipalID)
	if err != nil {
		return Decision{}, fmt.Errorf("load policies: %w", err)
	}
	memberships, err := s.store.MembershipsOf(ctx, req.PrincipalType, req.PrincipalID)
	if err != nil {
		return Decision{}, fmt.Errorf("load memberships: %w", err)
	}

	// No policies at all -> Cedar's default is deny, so short-circuit.
	if len(policies) == 0 {
		return Decision{Allowed: false, Reason: "no policies apply to principal"}, nil
	}

	// 2. Turn the structured policies into one Cedar document and parse it.
	policySet, err := buildPolicySet(policies)
	if err != nil {
		return Decision{}, fmt.Errorf("build policy set: %w", err)
	}

	// 3. Build the entity graph: the principal plus its role/group parents.
	//    Cedar needs this to resolve `principal in Role::"..."` checks.
	entities := buildEntities(req, memberships)

	// 4. Build the Cedar request (principal · action · resource · context).
	cedarReq := cedar.Request{
		Principal: cedar.NewEntityUID(cedar.EntityType(req.PrincipalType), cedar.String(req.PrincipalID)),
		Action:    cedar.NewEntityUID("Action", cedar.String(req.Action)),
		Resource:  cedar.NewEntityUID(cedar.EntityType(req.ResourceType), cedar.String(req.ResourceID)),
		Context:   buildContext(req.Context),
	}

	// 5. Evaluate. Decision is cedar.Allow or cedar.Deny; Diagnostic explains why.
	decision, diag := policySet.IsAuthorized(entities, cedarReq)

	return Decision{
		Allowed: decision == cedar.Allow,
		Reason:  summarize(decision, diag),
	}, nil
}

/* -------------------------------------------------------------------------- */
/*                                  helpers                                   */
/* -------------------------------------------------------------------------- */

// buildPolicySet concatenates every policy's Cedar text into a single document
// and parses it. Cedar happily reads many permit/forbid statements at once.
func buildPolicySet(policies []sentinelengine_models.Policy) (*cedar.PolicySet, error) {
	var doc strings.Builder
	for _, p := range policies {
		doc.WriteString(p.ToCedar())
		doc.WriteString("\n\n")
	}
	return cedar.NewPolicySetFromBytes("sentinel-policies", []byte(doc.String()))
}

// buildEntities describes the principal to Cedar, listing its roles/groups as
// "parents". That parentage is what makes a Role/Group policy apply to a User.
func buildEntities(req EvaluationRequest, memberships []Membership) cedar.EntityMap {
	entities := cedar.EntityMap{}

	principalUID := cedar.NewEntityUID(cedar.EntityType(req.PrincipalType), cedar.String(req.PrincipalID))

	parents := make([]cedar.EntityUID, 0, len(memberships))
	for _, m := range memberships {
		muid := cedar.NewEntityUID(cedar.EntityType(m.Type), cedar.String(m.ID))
		parents = append(parents, muid)
		// Each role/group must also exist as its own (parent-less) entity.
		entities[muid] = cedar.Entity{UID: muid}
	}

	entities[principalUID] = cedar.Entity{
		UID:     principalUID,
		Parents: cedar.NewEntityUIDSet(parents...),
	}

	// Describe the resource, and if it has a parent (e.g. its Property), record
	// that parentage so `resource is Restaurant in Property::"..."` can match.
	resourceUID := cedar.NewEntityUID(cedar.EntityType(req.ResourceType), cedar.String(req.ResourceID))
	var resourceParents []cedar.EntityUID
	if req.ResourceParentType != "" && req.ResourceParentID != "" {
		puid := cedar.NewEntityUID(cedar.EntityType(req.ResourceParentType), cedar.String(req.ResourceParentID))
		resourceParents = append(resourceParents, puid)
		entities[puid] = cedar.Entity{UID: puid}
	}
	entities[resourceUID] = cedar.Entity{
		UID:     resourceUID,
		Parents: cedar.NewEntityUIDSet(resourceParents...),
	}

	return entities
}

// buildContext converts the optional Context map into a Cedar record so policies
// can reference `context.<key>`. Only string/bool/int are supported.
func buildContext(ctx map[string]any) cedar.Record {
	if len(ctx) == 0 {
		return cedar.Record{}
	}
	rec := cedar.RecordMap{}
	for k, v := range ctx {
		switch val := v.(type) {
		case string:
			rec[cedar.String(k)] = cedar.String(val)
		case bool:
			rec[cedar.String(k)] = cedar.Boolean(val)
		case int:
			rec[cedar.String(k)] = cedar.Long(val)
		case int64:
			rec[cedar.String(k)] = cedar.Long(val)
		}
	}
	return cedar.NewRecord(rec)
}

// summarize turns a Cedar decision + diagnostic into a short log-friendly string.
func summarize(decision cedar.Decision, diag cedar.Diagnostic) string {
	if decision == cedar.Allow {
		return fmt.Sprintf("allowed by %d policy/policies", len(diag.Reasons))
	}
	if len(diag.Errors) > 0 {
		return fmt.Sprintf("denied (with %d evaluation error(s))", len(diag.Errors))
	}
	return "denied (no permit matched, or a forbid overrode it)"
}
