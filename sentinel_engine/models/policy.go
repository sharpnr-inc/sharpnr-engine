package sentinelengine_models

import (
	"fmt"
	"strings"
)

// This file defines how a SharpNR authorization rule is stored, and how it is
// turned into Cedar policy text that the engine can evaluate.
//
// You can describe a policy in TWO ways:
//
//  1. Structured fields (Effect / PrincipalType / Actions / ResourceType).
//     Good for the common case — "engineers may read & list Documents". The
//     engine calls ToCedar() to generate the Cedar text for you.
//
//  2. RawCedar — a hand-written Cedar string. Use this when you need a
//     condition the structured fields can't express (ABAC, `when {}` clauses,
//     specific resource IDs, etc.). When RawCedar is set, it is used verbatim
//     and the structured fields are ignored.

/* ----------------------------- small enums ------------------------------- */

// Effect is whether a policy grants or denies. In Cedar a `forbid` always beats
// a `permit`, and anything not explicitly permitted is denied by default.
type Effect string

const (
	Permit Effect = "permit"
	Forbid Effect = "forbid"
)

// PrincipalKind is the kind of entity a policy targets. Policies are normally
// attached to a Role or a Group; individual Users inherit them by membership.
type PrincipalKind string

const (
	PrincipalUser  PrincipalKind = "User"
	PrincipalRole  PrincipalKind = "Role"
	PrincipalGroup PrincipalKind = "Group"
)

/* ------------------------------- the model ------------------------------- */

// Policy is one authorization rule.
type Policy struct {
	// ID is a human-readable name, e.g. "engineer-read-docs". It only appears
	// as a comment in the generated Cedar; it is not used for matching.
	ID string

	// RawCedar, when non-empty, is used exactly as written (escape hatch).
	RawCedar string

	// Structured fields — used only when RawCedar is empty.
	Effect        Effect        // Permit or Forbid
	PrincipalType PrincipalKind // User / Role / Group
	PrincipalID   string        // e.g. "engineer", "billing-team", or a user id
	Actions       []string      // e.g. ["read", "list"]; empty == ALL actions
	ResourceType  string        // e.g. "Document", "Restaurant"; "" or "*" == ANY type

	// Optional resource scope: restrict the rule to resources that live UNDER a
	// given parent entity, e.g. ResourceInType "Property", ResourceInID "<uuid>".
	// This is how per-property access is expressed. Leave empty for no scope.
	ResourceInType string
	ResourceInID   string
}

/* ------------------------------- ToCedar --------------------------------- */

// ToCedar returns the Cedar text for this policy.
//
// Example — this Policy:
//
//	Policy{
//	    ID: "engineer-read-docs", Effect: Permit,
//	    PrincipalType: PrincipalRole, PrincipalID: "engineer",
//	    Actions: []string{"read", "list"}, ResourceType: "Document",
//	}
//
// becomes:
//
//	// Policy: engineer-read-docs
//	permit (
//	  principal in Role::"engineer",
//	  action in [Action::"read", Action::"list"],
//	  resource is Document
//	);
//
// Note: we use `in` (not `==`) for the principal. The request's principal is
// always the *User*; `in` means "the user is, or is a member of" the Role/Group,
// so role/group policies match through membership. For a User-typed policy `in`
// still works because an entity is always "in" itself.
func (p Policy) ToCedar() string {
	if strings.TrimSpace(p.RawCedar) != "" {
		return p.RawCedar
	}

	effect := p.Effect
	if effect == "" {
		effect = Permit
	}

	principalClause := fmt.Sprintf("principal in %s::\"%s\"", p.PrincipalType, p.PrincipalID)

	return fmt.Sprintf(
		"// Policy: %s\n%s (\n  %s,\n  %s,\n  %s\n);",
		p.ID,
		effect,
		principalClause,
		p.actionClause(),
		p.resourceClause(),
	)
}

// actionClause renders the action part. An empty Actions list means "any
// action", written in Cedar as a bare `action`.
func (p Policy) actionClause() string {
	if len(p.Actions) == 0 {
		return "action"
	}
	quoted := make([]string, len(p.Actions))
	for i, a := range p.Actions {
		quoted[i] = fmt.Sprintf(`Action::"%s"`, a)
	}
	return fmt.Sprintf("action in [%s]", strings.Join(quoted, ", "))
}

// resourceClause renders the resource part, combining type and optional scope:
//
//	any type, no scope      -> resource
//	any type, under parent  -> resource in Property::"<id>"
//	typed,    no scope       -> resource is Restaurant
//	typed,    under parent   -> resource is Restaurant in Property::"<id>"
func (p Policy) resourceClause() string {
	anyType := p.ResourceType == "" || p.ResourceType == "*"
	hasScope := p.ResourceInType != "" && p.ResourceInID != ""
	scope := fmt.Sprintf("%s::\"%s\"", p.ResourceInType, p.ResourceInID)

	switch {
	case anyType && hasScope:
		return fmt.Sprintf("resource in %s", scope)
	case anyType:
		return "resource"
	case hasScope:
		return fmt.Sprintf("resource is %s in %s", p.ResourceType, scope)
	default:
		return fmt.Sprintf("resource is %s", p.ResourceType)
	}
}
